package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCount200 = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_200",
		Help: "Total number of HTTP requests successfully served.",
	}, []string{"path", "method", "status"})
	RequestCount400 = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_400",
		Help: "Total number of HTTP requests not successfully served because of a client error.",
	}, []string{"path", "method", "status"})
	RequestCount500 = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_500",
		Help: "Total number of HTTP requests not successfully served because of a server error.",
	}, []string{"path", "method", "status"})
	RequestCountTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests made.",
	}, []string{})
	RequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration",
		Help:    "HTTP request latencies in seconds.",
		Buckets: prometheus.DefBuckets,
	}, []string{"path", "method"})
)

func init() {
	prometheus.MustRegister(RequestCount200, RequestCount400, RequestCount500, RequestCountTotal, RequestDuration)
}
