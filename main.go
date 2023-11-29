package main

import (
	"fmt"
	"kubetronix/probe/metrics"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		metrics.GetMetrics()
		promhttp.Handler().ServeHTTP(w, r)
	})
	port := ":9100"
	fmt.Printf("Exporter HTTP running %s port...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("HTTP Server not started: %v\n", err)
		os.Exit(1)
	}
}
