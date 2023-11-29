package counter

import (
	"encoding/json"
	"fmt"
	"os"
)

func LangCounters() {
	//prometheus.MustRegister(langMetric)

	apiEndpoint = os.Getenv("API_ENDPOINT")
	if apiEndpoint == "" {
		fmt.Println("API_ENDPOINT çevresel değişkeni belirtilmedi.")
		os.Exit(1)
	}
	// API'den veri çekme işlemi burada yapılır
	data, err := fetchData(apiEndpoint + "/api/dashboard/packagesPerLanguage")
	if err != nil {
		fmt.Printf("API'den veri çekilemedi: %v\n", err)
		return
	}

	// JSON verisini parse et
	var packagesPerLanguage []map[string]interface{}
	err = json.Unmarshal(data, &packagesPerLanguage)
	if err != nil {
		fmt.Printf("JSON verisi parse edilemedi: %v\n", err)
		return
	}

	// Metrikleri güncelle
	langMetric.Reset()

	for _, v := range packagesPerLanguage {
		language, ok := v["language"].(string)
		if !ok {
			fmt.Println("Zafiyet packagesPerLanguage değeri okunamadı.")
			continue
		}

		perLanguage, ok := v["count"].(float64)
		if !ok {
			fmt.Println("Zafiyet license değeri okunamadı.")
			continue
		}

		// Metrikleri güncelle
		langMetric.WithLabelValues(language).Set(perLanguage)
	}
}
