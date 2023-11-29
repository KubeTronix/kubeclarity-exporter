package counter

import (
	"encoding/json"
	"fmt"
	"os"
)

func VulnCounters() {

	apiEndpoint = os.Getenv("API_ENDPOINT")
	if apiEndpoint == "" {
		fmt.Println("API_ENDPOINT çevresel değişkeni belirtilmedi.")
		os.Exit(1)
	}
	// API'den veri çekme işlemi burada yapılır
	data, err := fetchData(apiEndpoint + "/api/dashboard/vulnerabilitiesWithFix")
	if err != nil {
		fmt.Printf("API'den veri çekilemedi: %v\n", err)
		return
	}

	// JSON verisini parse et
	var vulnerabilities []map[string]interface{}
	err = json.Unmarshal(data, &vulnerabilities)
	if err != nil {
		fmt.Printf("JSON verisi parse edilemedi: %v\n", err)
		return
	}

	// Metrikleri güncelle
	vulnFixMetric.Reset()
	vulnMetric.Reset()

	for _, v := range vulnerabilities {
		severity, ok := v["severity"].(string)
		if !ok {
			fmt.Println("Zafiyet severity değeri okunamadı.")
			continue
		}

		countWithFix, ok := v["countWithFix"].(float64)
		if !ok {
			fmt.Println("Zafiyet countWithFix değeri okunamadı.")
			continue
		}
		countTotal, ok := v["countTotal"].(float64)
		if !ok {
			fmt.Println("Zafiyet countTotal değeri okunamadı.")
			continue
		}

		// Metrikleri güncelle
		vulnFixMetric.WithLabelValues(severity).Set(countWithFix)
		vulnMetric.WithLabelValues(severity).Set(countTotal)
	}
}
