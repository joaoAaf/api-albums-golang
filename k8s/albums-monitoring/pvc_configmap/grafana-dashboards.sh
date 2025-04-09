#!/bin/sh
DASHBOARD_PATH=$1
kubectl create configmap grafana-dashboards \
  --from-file="${DASHBOARD_PATH}/dashboards.yml" \
  --from-file="${DASHBOARD_PATH}/albums_http_requests.json" \
  --from-file="${DASHBOARD_PATH}/metrics_go.json" \
  -n albums-monitoring