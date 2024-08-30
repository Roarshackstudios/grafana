# Use the official Grafana image from Docker Hub
FROM grafana/grafana:latest

# Set environment variables for Grafana admin credentials
ENV GF_SECURITY_ADMIN_USER=admin
ENV GF_SECURITY_ADMIN_PASSWORD=password

# Expose port 3000 (default port for Grafana)
EXPOSE 3000
