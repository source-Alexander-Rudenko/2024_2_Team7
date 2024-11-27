package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response times for handler in seconds",
			Buckets: prometheus.DefBuckets, // Например, 0.005, 0.01, 0.025, 0.05, ..., 10
		},
		[]string{"path", "method", "service", "status_code"},
	)

	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of requests",
		},
		[]string{"path", "method", "service", "status_code"},
	)

	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of error requests",
		},
		[]string{"path", "method", "service", "status_code"},
	)
)

func InitMetrics() {
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(ErrorCount)
}

func RegisterMetrics() {
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(ErrorCount)
}

func Listen(host string) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(host, mux)
}