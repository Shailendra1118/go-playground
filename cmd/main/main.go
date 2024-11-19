package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fmt.Printf("Hello, Go!\n")

	// Create a custom Prometheus counter
	httpRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests processed.",
		},
		[]string{"path"},
	)
	// Register the counter with Prometheus's default registry
	prometheus.MustRegister(httpRequestsTotal)


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Increment the counter for the root URL path
		httpRequestsTotal.WithLabelValues("/").Inc()
		fmt.Fprint(w, "Hello, Go Baby!")
	})
	// Expose the metrics endpoint at /metrics
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":8222", nil)
	if err != nil {
		fmt.Printf("error while starting server: %v", err)
	}
}
