docker run -d \
-p 3000:3000 \
--name=grafanaserver \
-e "GF_SECURITY_ADMIN_USER=admin" \
-e "GF_SECURITY_ADMIN_PASSWORD=password" \
grafana/grafana:latest
