// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     public/app/plugins/gen.go
// Using jennies:
//     PluginGoTypesJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package dataquery

// Defines values for CloudWatchQueryMode.
const (
	CloudWatchQueryModeAnnotations CloudWatchQueryMode = "Annotations"
	CloudWatchQueryModeLogs        CloudWatchQueryMode = "Logs"
	CloudWatchQueryModeMetrics     CloudWatchQueryMode = "Metrics"
)

// Defines values for MetricEditorMode.
const (
	MetricEditorModeN0 MetricEditorMode = 0
	MetricEditorModeN1 MetricEditorMode = 1
)

// Defines values for MetricQueryType.
const (
	MetricQueryTypeN0 MetricQueryType = 0
	MetricQueryTypeN1 MetricQueryType = 1
)

// Defines values for QueryEditorArrayExpressionType.
const (
	QueryEditorArrayExpressionTypeAnd QueryEditorArrayExpressionType = "and"
	QueryEditorArrayExpressionTypeOr  QueryEditorArrayExpressionType = "or"
)

// Defines values for QueryEditorExpressionType.
const (
	QueryEditorExpressionTypeAnd               QueryEditorExpressionType = "and"
	QueryEditorExpressionTypeFunction          QueryEditorExpressionType = "function"
	QueryEditorExpressionTypeFunctionParameter QueryEditorExpressionType = "functionParameter"
	QueryEditorExpressionTypeGroupBy           QueryEditorExpressionType = "groupBy"
	QueryEditorExpressionTypeOperator          QueryEditorExpressionType = "operator"
	QueryEditorExpressionTypeOr                QueryEditorExpressionType = "or"
	QueryEditorExpressionTypeProperty          QueryEditorExpressionType = "property"
)

// Defines values for QueryEditorFunctionExpressionType.
const (
	QueryEditorFunctionExpressionTypeAnd               QueryEditorFunctionExpressionType = "and"
	QueryEditorFunctionExpressionTypeFunction          QueryEditorFunctionExpressionType = "function"
	QueryEditorFunctionExpressionTypeFunctionParameter QueryEditorFunctionExpressionType = "functionParameter"
	QueryEditorFunctionExpressionTypeGroupBy           QueryEditorFunctionExpressionType = "groupBy"
	QueryEditorFunctionExpressionTypeOperator          QueryEditorFunctionExpressionType = "operator"
	QueryEditorFunctionExpressionTypeOr                QueryEditorFunctionExpressionType = "or"
	QueryEditorFunctionExpressionTypeProperty          QueryEditorFunctionExpressionType = "property"
)

// Defines values for QueryEditorFunctionParameterExpressionType.
const (
	QueryEditorFunctionParameterExpressionTypeAnd               QueryEditorFunctionParameterExpressionType = "and"
	QueryEditorFunctionParameterExpressionTypeFunction          QueryEditorFunctionParameterExpressionType = "function"
	QueryEditorFunctionParameterExpressionTypeFunctionParameter QueryEditorFunctionParameterExpressionType = "functionParameter"
	QueryEditorFunctionParameterExpressionTypeGroupBy           QueryEditorFunctionParameterExpressionType = "groupBy"
	QueryEditorFunctionParameterExpressionTypeOperator          QueryEditorFunctionParameterExpressionType = "operator"
	QueryEditorFunctionParameterExpressionTypeOr                QueryEditorFunctionParameterExpressionType = "or"
	QueryEditorFunctionParameterExpressionTypeProperty          QueryEditorFunctionParameterExpressionType = "property"
)

// Defines values for QueryEditorGroupByExpressionType.
const (
	QueryEditorGroupByExpressionTypeAnd               QueryEditorGroupByExpressionType = "and"
	QueryEditorGroupByExpressionTypeFunction          QueryEditorGroupByExpressionType = "function"
	QueryEditorGroupByExpressionTypeFunctionParameter QueryEditorGroupByExpressionType = "functionParameter"
	QueryEditorGroupByExpressionTypeGroupBy           QueryEditorGroupByExpressionType = "groupBy"
	QueryEditorGroupByExpressionTypeOperator          QueryEditorGroupByExpressionType = "operator"
	QueryEditorGroupByExpressionTypeOr                QueryEditorGroupByExpressionType = "or"
	QueryEditorGroupByExpressionTypeProperty          QueryEditorGroupByExpressionType = "property"
)

// Defines values for QueryEditorOperatorExpressionType.
const (
	QueryEditorOperatorExpressionTypeAnd               QueryEditorOperatorExpressionType = "and"
	QueryEditorOperatorExpressionTypeFunction          QueryEditorOperatorExpressionType = "function"
	QueryEditorOperatorExpressionTypeFunctionParameter QueryEditorOperatorExpressionType = "functionParameter"
	QueryEditorOperatorExpressionTypeGroupBy           QueryEditorOperatorExpressionType = "groupBy"
	QueryEditorOperatorExpressionTypeOperator          QueryEditorOperatorExpressionType = "operator"
	QueryEditorOperatorExpressionTypeOr                QueryEditorOperatorExpressionType = "or"
	QueryEditorOperatorExpressionTypeProperty          QueryEditorOperatorExpressionType = "property"
)

// Defines values for QueryEditorPropertyExpressionType.
const (
	QueryEditorPropertyExpressionTypeAnd               QueryEditorPropertyExpressionType = "and"
	QueryEditorPropertyExpressionTypeFunction          QueryEditorPropertyExpressionType = "function"
	QueryEditorPropertyExpressionTypeFunctionParameter QueryEditorPropertyExpressionType = "functionParameter"
	QueryEditorPropertyExpressionTypeGroupBy           QueryEditorPropertyExpressionType = "groupBy"
	QueryEditorPropertyExpressionTypeOperator          QueryEditorPropertyExpressionType = "operator"
	QueryEditorPropertyExpressionTypeOr                QueryEditorPropertyExpressionType = "or"
	QueryEditorPropertyExpressionTypeProperty          QueryEditorPropertyExpressionType = "property"
)

// Defines values for QueryEditorPropertyType.
const (
	QueryEditorPropertyTypeString QueryEditorPropertyType = "string"
)

// Shape of a CloudWatch Annotation query
//
// TS type is CloudWatchDefaultQuery = Omit<CloudWatchLogsQuery, 'queryMode'> & CloudWatchMetricsQuery, declared in veneer
// #CloudWatchDefaultQuery: #CloudWatchLogsQuery & #CloudWatchMetricsQuery @cuetsy(kind="type")
type CloudWatchAnnotationQuery struct {
	// The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
	AccountId *string `json:"accountId,omitempty"`

	// Use this parameter to filter the results of the operation to only those alarms
	// that use a certain alarm action. For example, you could specify the ARN of
	// an SNS topic to find all alarms that send notifications to that topic.
	// e.g. `arn:aws:sns:us-east-1:123456789012:my-app-` would match `arn:aws:sns:us-east-1:123456789012:my-app-action`
	// but not match `arn:aws:sns:us-east-1:123456789012:your-app-action`
	ActionPrefix *string `json:"actionPrefix,omitempty"`

	// An alarm name prefix. If you specify this parameter, you receive information
	// about all alarms that have names that start with this prefix.
	// e.g. `my-team-service-` would match `my-team-service-high-cpu` but not match `your-team-service-high-cpu`
	AlarmNamePrefix *string `json:"alarmNamePrefix,omitempty"`

	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *any `json:"datasource,omitempty"`

	// A name/value pair that is part of the identity of a metric. For example, you can get statistics for a specific EC2 instance by specifying the InstanceId dimension when you search for metrics.
	Dimensions *Dimensions `json:"dimensions,omitempty"`

	// Hide true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`

	// Only show metrics that exactly match all defined dimension names.
	MatchExact *bool `json:"matchExact,omitempty"`

	// Name of the metric
	MetricName *string `json:"metricName,omitempty"`

	// A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
	Namespace *string `json:"namespace,omitempty"`

	// The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
	Period *string `json:"period,omitempty"`

	// Enable matching on the prefix of the action name or alarm name, specify the prefixes with actionPrefix and/or alarmNamePrefix
	PrefixMatching *bool                `json:"prefixMatching,omitempty"`
	QueryMode      *CloudWatchQueryMode `json:"queryMode,omitempty"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId *string `json:"refId,omitempty"`

	// AWS region to query for the metric
	Region *string `json:"region,omitempty"`

	// Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
	Statistic *string `json:"statistic,omitempty"`

	// @deprecated use statistic
	Statistics []string `json:"statistics,omitempty"`
}

// CloudWatchDataQuery defines model for CloudWatchDataQuery.
type CloudWatchDataQuery = map[string]any

// Shape of a CloudWatch Logs query
type CloudWatchLogsQuery struct {
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *any `json:"datasource,omitempty"`

	// The CloudWatch Logs Insights query to execute
	Expression *string `json:"expression,omitempty"`

	// Hide true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool   `json:"hide,omitempty"`
	Id   *string `json:"id,omitempty"`

	// @deprecated use logGroups
	LogGroupNames []string `json:"logGroupNames,omitempty"`

	// Log groups to query
	LogGroups []LogGroup           `json:"logGroups,omitempty"`
	QueryMode *CloudWatchQueryMode `json:"queryMode,omitempty"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId *string `json:"refId,omitempty"`

	// AWS region to query for the logs
	Region *string `json:"region,omitempty"`

	// Fields to group the results by, this field is automatically populated whenever the query is updated
	StatsGroups []string `json:"statsGroups,omitempty"`
}

// Shape of a CloudWatch Metrics query
type CloudWatchMetricsQuery struct {
	// The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
	AccountId *string `json:"accountId,omitempty"`

	// Deprecated: use label
	// @deprecated use label
	Alias *string `json:"alias,omitempty"`

	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *any `json:"datasource,omitempty"`

	// A name/value pair that is part of the identity of a metric. For example, you can get statistics for a specific EC2 instance by specifying the InstanceId dimension when you search for metrics.
	Dimensions *Dimensions `json:"dimensions,omitempty"`

	// Math expression query
	Expression *string `json:"expression,omitempty"`

	// Hide true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`

	// ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.
	Id *string `json:"id,omitempty"`

	// Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.
	Label *string `json:"label,omitempty"`

	// Only show metrics that exactly match all defined dimension names.
	MatchExact       *bool             `json:"matchExact,omitempty"`
	MetricEditorMode *MetricEditorMode `json:"metricEditorMode,omitempty"`

	// Name of the metric
	MetricName      *string          `json:"metricName,omitempty"`
	MetricQueryType *MetricQueryType `json:"metricQueryType,omitempty"`

	// A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
	Namespace *string `json:"namespace,omitempty"`

	// The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
	Period    *string              `json:"period,omitempty"`
	QueryMode *CloudWatchQueryMode `json:"queryMode,omitempty"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId *string `json:"refId,omitempty"`

	// AWS region to query for the metric
	Region *string        `json:"region,omitempty"`
	Sql    *SQLExpression `json:"sql,omitempty"`

	// When the metric query type is `metricQueryType` is set to `Query`, this field is used to specify the query string.
	SqlExpression *string `json:"sqlExpression,omitempty"`

	// Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
	Statistic *string `json:"statistic,omitempty"`

	// @deprecated use statistic
	Statistics []string `json:"statistics,omitempty"`
}

// CloudWatchQueryMode defines model for CloudWatchQueryMode.
type CloudWatchQueryMode string

// These are the common properties available to all queries in all datasources.
// Specific implementations will *extend* this interface, adding the required
// properties for the given context.
type DataQuery struct {
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *any `json:"datasource,omitempty"`

	// Hide true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
}

// A name/value pair that is part of the identity of a metric. For example, you can get statistics for a specific EC2 instance by specifying the InstanceId dimension when you search for metrics.
type Dimensions map[string]any

// LogGroup defines model for LogGroup.
type LogGroup struct {
	// AccountId of the log group
	AccountId *string `json:"accountId,omitempty"`

	// Label of the log group
	AccountLabel *string `json:"accountLabel,omitempty"`

	// ARN of the log group
	Arn string `json:"arn"`

	// Name of the log group
	Name string `json:"name"`
}

// MetricEditorMode defines model for MetricEditorMode.
type MetricEditorMode int

// MetricQueryType defines model for MetricQueryType.
type MetricQueryType int

// MetricStat defines model for MetricStat.
type MetricStat struct {
	// The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
	AccountId *string `json:"accountId,omitempty"`

	// A name/value pair that is part of the identity of a metric. For example, you can get statistics for a specific EC2 instance by specifying the InstanceId dimension when you search for metrics.
	Dimensions *Dimensions `json:"dimensions,omitempty"`

	// Only show metrics that exactly match all defined dimension names.
	MatchExact *bool `json:"matchExact,omitempty"`

	// Name of the metric
	MetricName *string `json:"metricName,omitempty"`

	// A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
	Namespace string `json:"namespace"`

	// The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
	Period *string `json:"period,omitempty"`

	// AWS region to query for the metric
	Region string `json:"region"`

	// Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
	Statistic *string `json:"statistic,omitempty"`

	// @deprecated use statistic
	Statistics []string `json:"statistics,omitempty"`
}

// QueryEditorArrayExpression defines model for QueryEditorArrayExpression.
type QueryEditorArrayExpression struct {
	Expressions []any                          `json:"expressions"`
	Type        QueryEditorArrayExpressionType `json:"type"`
}

// QueryEditorArrayExpressionType defines model for QueryEditorArrayExpression.Type.
type QueryEditorArrayExpressionType string

// QueryEditorExpressionType defines model for QueryEditorExpressionType.
type QueryEditorExpressionType string

// QueryEditorFunctionExpression defines model for QueryEditorFunctionExpression.
type QueryEditorFunctionExpression struct {
	Name       *string                                  `json:"name,omitempty"`
	Parameters []QueryEditorFunctionParameterExpression `json:"parameters,omitempty"`
	Type       QueryEditorFunctionExpressionType        `json:"type"`
}

// QueryEditorFunctionExpressionType defines model for QueryEditorFunctionExpression.Type.
type QueryEditorFunctionExpressionType string

// QueryEditorFunctionParameterExpression defines model for QueryEditorFunctionParameterExpression.
type QueryEditorFunctionParameterExpression struct {
	Name *string                                    `json:"name,omitempty"`
	Type QueryEditorFunctionParameterExpressionType `json:"type"`
}

// QueryEditorFunctionParameterExpressionType defines model for QueryEditorFunctionParameterExpression.Type.
type QueryEditorFunctionParameterExpressionType string

// QueryEditorGroupByExpression defines model for QueryEditorGroupByExpression.
type QueryEditorGroupByExpression struct {
	Property QueryEditorProperty              `json:"property"`
	Type     QueryEditorGroupByExpressionType `json:"type"`
}

// QueryEditorGroupByExpressionType defines model for QueryEditorGroupByExpression.Type.
type QueryEditorGroupByExpressionType string

// TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
type QueryEditorOperator struct {
	Name  *string `json:"name,omitempty"`
	Value *any    `json:"value,omitempty"`
}

// QueryEditorOperatorExpression defines model for QueryEditorOperatorExpression.
type QueryEditorOperatorExpression struct {
	// TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
	Operator QueryEditorOperator               `json:"operator"`
	Property QueryEditorProperty               `json:"property"`
	Type     QueryEditorOperatorExpressionType `json:"type"`
}

// QueryEditorOperatorExpressionType defines model for QueryEditorOperatorExpression.Type.
type QueryEditorOperatorExpressionType string

// QueryEditorProperty defines model for QueryEditorProperty.
type QueryEditorProperty struct {
	Name *string                 `json:"name,omitempty"`
	Type QueryEditorPropertyType `json:"type"`
}

// QueryEditorPropertyExpression defines model for QueryEditorPropertyExpression.
type QueryEditorPropertyExpression struct {
	Property QueryEditorProperty               `json:"property"`
	Type     QueryEditorPropertyExpressionType `json:"type"`
}

// QueryEditorPropertyExpressionType defines model for QueryEditorPropertyExpression.Type.
type QueryEditorPropertyExpressionType string

// QueryEditorPropertyType defines model for QueryEditorPropertyType.
type QueryEditorPropertyType string

// SQLExpression defines model for SQLExpression.
type SQLExpression struct {
	// FROM part of the SQL expression
	From    *any                        `json:"from,omitempty"`
	GroupBy *QueryEditorArrayExpression `json:"groupBy,omitempty"`

	// LIMIT part of the SQL expression
	Limit   *int64                         `json:"limit,omitempty"`
	OrderBy *QueryEditorFunctionExpression `json:"orderBy,omitempty"`

	// The sort order of the SQL expression, `ASC` or `DESC`
	OrderByDirection *string                        `json:"orderByDirection,omitempty"`
	Select           *QueryEditorFunctionExpression `json:"select,omitempty"`
	Where            *QueryEditorArrayExpression    `json:"where,omitempty"`
}
