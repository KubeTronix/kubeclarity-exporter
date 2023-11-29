package counter

import (
	"encoding/json"
	"fmt"
	"os"
)

func LicenseCounters() {

	apiEndpoint = os.Getenv("API_ENDPOINT")
	if apiEndpoint == "" {
		fmt.Println("API_ENDPOINT çevresel değişkeni belirtilmedi.")
		os.Exit(1)
	}
	// API'den veri çekme işlemi burada yapılır
	data, err := fetchData(apiEndpoint + "/api/dashboard/packagesPerLicense")
	if err != nil {
		fmt.Printf("API'den veri çekilemedi: %v\n", err)
		return
	}

	// JSON verisini parse et
	var packagesPerLicense []map[string]interface{}
	err = json.Unmarshal(data, &packagesPerLicense)
	if err != nil {
		fmt.Printf("JSON verisi parse edilemedi: %v\n", err)
		return
	}

	// Metrikleri güncelle
	licenseMetric.Reset()

	for _, v := range packagesPerLicense {
		license, ok := v["license"].(string)
		if !ok {
			fmt.Println("Zafiyet packagesPerLicense değeri okunamadı.")
			continue
		}

		perLicense, ok := v["count"].(float64)
		if !ok {
			fmt.Println("Zafiyet license değeri okunamadı.")
			continue
		}

		// Metrikleri güncelle
		licenseMetric.WithLabelValues(license).Set(perLicense)
	}
}
