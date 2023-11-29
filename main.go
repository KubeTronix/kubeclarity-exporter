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
	//http.Handle("/metrics", promhttp.Handler())
	port := ":9100"
	fmt.Printf("Exporter HTTP sunucusu %s portunda çalışıyor...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("HTTP sunucusu başlatılamadı: %v\n", err)
		os.Exit(1)
	}
}
