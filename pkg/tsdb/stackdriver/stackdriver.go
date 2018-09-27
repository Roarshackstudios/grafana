package stackdriver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context/ctxhttp"

	"github.com/grafana/grafana/pkg/api/pluginproxy"
	"github.com/grafana/grafana/pkg/cmd/grafana-cli/logger"
	"github.com/grafana/grafana/pkg/components/null"
	"github.com/grafana/grafana/pkg/components/simplejson"
	"github.com/grafana/grafana/pkg/log"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/plugins"
	"github.com/grafana/grafana/pkg/setting"
	"github.com/grafana/grafana/pkg/tsdb"
	"github.com/opentracing/opentracing-go"
)

var (
	slog                  log.Logger
	legendKeyFormat       *regexp.Regexp
	longMetricNameFormat  *regexp.Regexp
	shortMetricNameFormat *regexp.Regexp
)

// StackdriverExecutor executes queries for the Stackdriver datasource
type StackdriverExecutor struct {
	httpClient *http.Client
	dsInfo     *models.DataSource
}

// NewStackdriverExecutor initializes a http client
func NewStackdriverExecutor(dsInfo *models.DataSource) (tsdb.TsdbQueryEndpoint, error) {
	httpClient, err := dsInfo.GetHttpClient()
	if err != nil {
		return nil, err
	}

	return &StackdriverExecutor{
		httpClient: httpClient,
		dsInfo:     dsInfo,
	}, nil
}

func init() {
	slog = log.New("tsdb.stackdriver")
	tsdb.RegisterTsdbQueryEndpoint("stackdriver", NewStackdriverExecutor)
	legendKeyFormat = regexp.MustCompile(`\{\{\s*(.+?)\s*\}\}`)
	longMetricNameFormat = regexp.MustCompile(`([\w\d_]+)\.googleapis\.com/([\w\d_]+)/(.+)`)
	shortMetricNameFormat = regexp.MustCompile(`([\w\d_]+)\.googleapis\.com/(.+)`)
}

// Query takes in the frontend queries, parses them into the Stackdriver query format
// executes the queries against the Stackdriver API and parses the response into
// the time series or table format
func (e *StackdriverExecutor) Query(ctx context.Context, dsInfo *models.DataSource, tsdbQuery *tsdb.TsdbQuery) (*tsdb.Response, error) {
	var result *tsdb.Response
	var err error
	queryType := tsdbQuery.Queries[0].Model.Get("type").MustString("")

	switch queryType {
	case "annotationQuery":
		result, err = e.executeAnnotationQuery(ctx, tsdbQuery)
	case "timeSeriesQuery":
		fallthrough
	default:
		result, err = e.executeTimeSeriesQuery(ctx, tsdbQuery)
	}

	return result, err
}

func (e *StackdriverExecutor) executeTimeSeriesQuery(ctx context.Context, tsdbQuery *tsdb.TsdbQuery) (*tsdb.Response, error) {
	result := &tsdb.Response{
		Results: make(map[string]*tsdb.QueryResult),
	}

	queries, err := e.buildQueries(tsdbQuery)
	if err != nil {
		return nil, err
	}

	for _, query := range queries {
		queryRes, err := e.executeQuery(ctx, query, tsdbQuery)
		if err != nil {
			return nil, err
		}
		result.Results[query.RefID] = queryRes
	}

	return result, nil
}

func (e *StackdriverExecutor) buildQueries(tsdbQuery *tsdb.TsdbQuery) ([]*StackdriverQuery, error) {
	stackdriverQueries := []*StackdriverQuery{}

	startTime, err := tsdbQuery.TimeRange.ParseFrom()
	if err != nil {
		return nil, err
	}

	endTime, err := tsdbQuery.TimeRange.ParseTo()
	if err != nil {
		return nil, err
	}

	durationSeconds := int(endTime.Sub(startTime).Seconds())

	for _, query := range tsdbQuery.Queries {
		var target string

		metricType := query.Model.Get("metricType").MustString()
		filterParts := query.Model.Get("filters").MustArray()

		params := url.Values{}
		params.Add("interval.startTime", startTime.UTC().Format(time.RFC3339))
		params.Add("interval.endTime", endTime.UTC().Format(time.RFC3339))
		params.Add("filter", buildFilterString(metricType, filterParts))
		params.Add("view", query.Model.Get("view").MustString())
		setAggParams(&params, query, durationSeconds)

		target = params.Encode()

		if setting.Env == setting.DEV {
			slog.Debug("Stackdriver request", "params", params)
		}

		groupBys := query.Model.Get("groupBys").MustArray()
		groupBysAsStrings := make([]string, 0)
		for _, groupBy := range groupBys {
			groupBysAsStrings = append(groupBysAsStrings, groupBy.(string))
		}

		aliasBy := query.Model.Get("aliasBy").MustString()

		stackdriverQueries = append(stackdriverQueries, &StackdriverQuery{
			Target:   target,
			Params:   params,
			RefID:    query.RefId,
			GroupBys: groupBysAsStrings,
			AliasBy:  aliasBy,
		})
	}

	return stackdriverQueries, nil
}

func buildFilterString(metricType string, filterParts []interface{}) string {
	filterString := ""
	for i, part := range filterParts {
		mod := i % 4
		if part == "AND" {
			filterString += " "
		} else if mod == 2 {
			filterString += fmt.Sprintf(`"%s"`, part)
		} else {
			filterString += part.(string)
		}
	}
	return strings.Trim(fmt.Sprintf(`metric.type="%s" %s`, metricType, filterString), " ")
}

func setAggParams(params *url.Values, query *tsdb.Query, durationSeconds int) {
	primaryAggregation := query.Model.Get("primaryAggregation").MustString()
	perSeriesAligner := query.Model.Get("perSeriesAligner").MustString()
	alignmentPeriod := query.Model.Get("alignmentPeriod").MustString()

	if primaryAggregation == "" {
		primaryAggregation = "REDUCE_NONE"
	}

	if perSeriesAligner == "" {
		perSeriesAligner = "ALIGN_MEAN"
	}

	if alignmentPeriod == "grafana-auto" || alignmentPeriod == "" {
		alignmentPeriodValue := int(math.Max(float64(query.IntervalMs)/1000, 60.0))
		alignmentPeriod = "+" + strconv.Itoa(alignmentPeriodValue) + "s"
	}

	if alignmentPeriod == "stackdriver-auto" {
		alignmentPeriodValue := int(math.Max(float64(durationSeconds), 60.0))
		logger.Info("alignmentPeriodValue", "alignmentPeriodValue", alignmentPeriodValue)
		if alignmentPeriodValue < 60*60*23 {
			alignmentPeriod = "+60s"
		} else if alignmentPeriodValue < 60*60*24*6 {
			alignmentPeriod = "+300s"
		} else {
			alignmentPeriod = "+3600s"
		}
	}

	re := regexp.MustCompile("[0-9]+")
	seconds, err := strconv.ParseInt(re.FindString(alignmentPeriod), 10, 64)
	if err != nil || seconds > 3600 {
		alignmentPeriod = "+3600s"
	}

	params.Add("aggregation.crossSeriesReducer", primaryAggregation)
	params.Add("aggregation.perSeriesAligner", perSeriesAligner)
	params.Add("aggregation.alignmentPeriod", alignmentPeriod)

	groupBys := query.Model.Get("groupBys").MustArray()
	if len(groupBys) > 0 {
		for i := 0; i < len(groupBys); i++ {
			params.Add("aggregation.groupByFields", groupBys[i].(string))
		}
	}
}

func (e *StackdriverExecutor) executeQuery(ctx context.Context, query *StackdriverQuery, tsdbQuery *tsdb.TsdbQuery) (*tsdb.QueryResult, error) {
	queryResult := &tsdb.QueryResult{Meta: simplejson.New(), RefId: query.RefID}

	req, err := e.createRequest(ctx, e.dsInfo)
	if err != nil {
		queryResult.Error = err
		return queryResult, nil
	}

	req.URL.RawQuery = query.Params.Encode()
	queryResult.Meta.Set("rawQuery", req.URL.RawQuery)
	alignmentPeriod, ok := req.URL.Query()["aggregation.alignmentPeriod"]

	if ok {
		re := regexp.MustCompile("[0-9]+")
		seconds, err := strconv.ParseInt(re.FindString(alignmentPeriod[0]), 10, 64)
		if err == nil {
			queryResult.Meta.Set("alignmentPeriod", seconds)
		}
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "stackdriver query")
	span.SetTag("target", query.Target)
	span.SetTag("from", tsdbQuery.TimeRange.From)
	span.SetTag("until", tsdbQuery.TimeRange.To)
	span.SetTag("datasource_id", e.dsInfo.Id)
	span.SetTag("org_id", e.dsInfo.OrgId)

	defer span.Finish()

	opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header))

	res, err := ctxhttp.Do(ctx, e.httpClient, req)
	if err != nil {
		queryResult.Error = err
		return queryResult, nil
	}

	data, err := e.unmarshalResponse(res)
	if err != nil {
		queryResult.Error = err
		return queryResult, nil
	}

	err = e.parseResponse(queryResult, data, query)
	if err != nil {
		queryResult.Error = err
		return queryResult, nil
	}

	return queryResult, nil
}

func (e *StackdriverExecutor) unmarshalResponse(res *http.Response) (StackdriverResponse, error) {
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return StackdriverResponse{}, err
	}

	if res.StatusCode/100 != 2 {
		slog.Error("Request failed", "status", res.Status, "body", string(body))
		return StackdriverResponse{}, fmt.Errorf(string(body))
	}

	var data StackdriverResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		slog.Error("Failed to unmarshal Stackdriver response", "error", err, "status", res.Status, "body", string(body))
		return StackdriverResponse{}, err
	}

	return data, nil
}

func (e *StackdriverExecutor) parseResponse(queryRes *tsdb.QueryResult, data StackdriverResponse, query *StackdriverQuery) error {
	metricLabels := make(map[string][]string)
	resourceLabels := make(map[string][]string)

	for _, series := range data.TimeSeries {
		points := make([]tsdb.TimePoint, 0)

		// reverse the order to be ascending
		for i := len(series.Points) - 1; i >= 0; i-- {
			point := series.Points[i]
			points = append(points, tsdb.NewTimePoint(null.FloatFrom(point.Value.DoubleValue), float64((point.Interval.EndTime).Unix())*1000))
		}

		defaultMetricName := series.Metric.Type

		for key, value := range series.Metric.Labels {
			if !containsLabel(metricLabels[key], value) {
				metricLabels[key] = append(metricLabels[key], value)
			}
			if len(query.GroupBys) == 0 || containsLabel(query.GroupBys, "metric.label."+key) {
				defaultMetricName += " " + value
			}
		}

		for key, value := range series.Resource.Labels {
			if !containsLabel(resourceLabels[key], value) {
				resourceLabels[key] = append(resourceLabels[key], value)
			}

			if containsLabel(query.GroupBys, "resource.label."+key) {
				defaultMetricName += " " + value
			}
		}

		metricName := formatLegendKeys(series.Metric.Type, defaultMetricName, series.Metric.Labels, series.Resource.Labels, query)

		queryRes.Series = append(queryRes.Series, &tsdb.TimeSeries{
			Name:   metricName,
			Points: points,
		})
	}

	queryRes.Meta.Set("resourceLabels", resourceLabels)
	queryRes.Meta.Set("metricLabels", metricLabels)
	queryRes.Meta.Set("groupBys", query.GroupBys)

	return nil
}

func containsLabel(labels []string, newLabel string) bool {
	for _, val := range labels {
		if val == newLabel {
			return true
		}
	}
	return false
}

func formatLegendKeys(metricType string, defaultMetricName string, metricLabels map[string]string, resourceLabels map[string]string, query *StackdriverQuery) string {
	if query.AliasBy == "" {
		return defaultMetricName
	}

	result := legendKeyFormat.ReplaceAllFunc([]byte(query.AliasBy), func(in []byte) []byte {
		metaPartName := strings.Replace(string(in), "{{", "", 1)
		metaPartName = strings.Replace(metaPartName, "}}", "", 1)
		metaPartName = strings.TrimSpace(metaPartName)

		if metaPartName == "metric.type" {
			return []byte(metricType)
		}

		metricPart := replaceWithMetricPart(metaPartName, metricType)

		if metricPart != nil {
			return metricPart
		}

		metaPartName = strings.Replace(metaPartName, "metric.label.", "", 1)

		if val, exists := metricLabels[metaPartName]; exists {
			return []byte(val)
		}

		metaPartName = strings.Replace(metaPartName, "resource.label.", "", 1)

		if val, exists := resourceLabels[metaPartName]; exists {
			return []byte(val)
		}

		return in
	})

	return string(result)
}

func replaceWithMetricPart(metaPartName string, metricType string) []byte {
	// https://cloud.google.com/monitoring/api/v3/metrics-details#label_names
	longMatches := longMetricNameFormat.FindStringSubmatch(metricType)
	shortMatches := shortMetricNameFormat.FindStringSubmatch(metricType)

	if metaPartName == "metric.name" {
		if len(longMatches) > 0 {
			return []byte(longMatches[3])
		} else if len(shortMatches) > 0 {
			return []byte(shortMatches[2])
		}
	}

	if metaPartName == "metric.category" {
		if len(longMatches) > 0 {
			return []byte(longMatches[2])
		}
	}

	if metaPartName == "metric.service" {
		if len(longMatches) > 0 {
			return []byte(longMatches[1])
		} else if len(shortMatches) > 0 {
			return []byte(shortMatches[1])
		}
	}

	return nil
}

func (e *StackdriverExecutor) createRequest(ctx context.Context, dsInfo *models.DataSource) (*http.Request, error) {
	u, _ := url.Parse(dsInfo.Url)
	u.Path = path.Join(u.Path, "render")

	req, err := http.NewRequest(http.MethodGet, "https://monitoring.googleapis.com/", nil)
	if err != nil {
		slog.Info("Failed to create request", "error", err)
		return nil, fmt.Errorf("Failed to create request. error: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", fmt.Sprintf("Grafana/%s", setting.BuildVersion))

	// find plugin
	plugin, ok := plugins.DataSources[dsInfo.Type]
	if !ok {
		return nil, errors.New("Unable to find datasource plugin Stackdriver")
	}
	projectName := dsInfo.JsonData.Get("defaultProject").MustString()
	proxyPass := fmt.Sprintf("stackdriver%s", "v3/projects/"+projectName+"/timeSeries")

	var stackdriverRoute *plugins.AppPluginRoute
	for _, route := range plugin.Routes {
		if route.Path == "stackdriver" {
			stackdriverRoute = route
			break
		}
	}

	pluginproxy.ApplyRoute(ctx, req, proxyPass, stackdriverRoute, dsInfo)

	return req, nil
}
