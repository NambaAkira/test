package service

import (
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

func PrometheusInterceptor() grpc.ServerOption {
	// Counter for requests.
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.EnableClientHandlingTimeHistogram()
	grpc_prometheus.EnableHandlingTimeSummary()

	// Register prometheus metrics.
	grpc_prometheus.Register(prometheus.DefaultRegisterer)

	// Create histogram for requests duration.
	histVecOpts := prometheus.HistogramOpts{
		Name:    "grpc_request_duration_seconds",
		Help:    "Histogram of grpc request duration in seconds.",
		Buckets: []float64{0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
	}

	histVec := prometheus.NewHistogramVec(histVecOpts, []string{"grpc_service", "grpc_method"})

	// Return an interceptor to register this histogram with Prometheus.
	return grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(prometheus.UnaryServerInterceptor(histVec)),
		grpc.StreamInterceptor(prometheus.StreamServerInterceptor(histVec))
}
