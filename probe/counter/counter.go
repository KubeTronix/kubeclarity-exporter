package counter

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetCounters() {

	// Çevresel değişkeni oku
	apiEndpoint = os.Getenv("API_ENDPOINT")
	if apiEndpoint == "" {
		fmt.Println("API_ENDPOINT çevresel değişkeni belirtilmedi.")
		os.Exit(1)
	}
	// API'den veri çekme işlemi burada yapılır
	data, err := fetchData(apiEndpoint + "/api/dashboard/counters")
	if err != nil {
		fmt.Printf("API'den veri çekilemedi: %v\n", err)
		return
	}

	// JSON verisini parse et
	var counters map[string]int
	err = json.Unmarshal(data, &counters)
	if err != nil {
		fmt.Printf("JSON verisi parse edilemedi: %v\n", err)
		return
	}

	// Metrikleri güncelle
	applicationsMetric.Reset()
	applicationsMetric.WithLabelValues().Set(float64(counters["applications"]))

	packagesMetric.Reset()
	packagesMetric.WithLabelValues().Set(float64(counters["packages"]))

	resourcesMetric.Reset()
	resourcesMetric.WithLabelValues().Set(float64(counters["resources"]))
}
