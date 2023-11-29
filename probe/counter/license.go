package counter

import (
	"encoding/json"
	"fmt"
	"os"
)

func LicenseCounters() {

	apiEndpoint = os.Getenv("API_ENDPOINT")
	if apiEndpoint == "" {
		fmt.Println("API_ENDPOINT not read env")
		os.Exit(1)
	}

	data, err := fetchData(apiEndpoint + "/api/dashboard/packagesPerLicense")
	if err != nil {
		fmt.Printf("Failed to retrieve data from API: %v\n", err)
		return
	}

	var packagesPerLicense []map[string]interface{}
	err = json.Unmarshal(data, &packagesPerLicense)
	if err != nil {
		fmt.Printf("JSON data could not be parsed: %v\n", err)
		return
	}
	licenseMetric.Reset()

	for _, v := range packagesPerLicense {
		license, ok := v["license"].(string)
		if !ok {
			fmt.Println("License value was not read in packagesPerLicense.")
			continue
		}

		perLicense, ok := v["count"].(float64)
		if !ok {
			fmt.Println("Count value was not read in packagesPerLicense.")
			continue
		}

		licenseMetric.WithLabelValues(license).Set(perLicense)
	}
}
