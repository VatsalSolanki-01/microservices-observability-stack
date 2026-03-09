package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"method", "endpoint"},
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	httpRequests.WithLabelValues(r.Method, "/health").Inc()
	w.Write([]byte("OK"))
}

func main() {

	prometheus.MustRegister(httpRequests)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", healthHandler)

	log.Println("User service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}