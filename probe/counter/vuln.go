package counter

import (
	"encoding/json"
	"fmt"
	"os"
)

func VulnCounters() {

	apiEndpoint = os.Getenv("API_ENDPOINT")
	if apiEndpoint == "" {
		fmt.Println("API_ENDPOINT not read env")
		os.Exit(1)
	}

	data, err := fetchData(apiEndpoint + "/api/dashboard/vulnerabilitiesWithFix")
	if err != nil {
		fmt.Printf("Failed to retrieve data from API: %v\n", err)
		return
	}

	var vulnerabilities []map[string]interface{}
	err = json.Unmarshal(data, &vulnerabilities)
	if err != nil {
		fmt.Printf("JSON data could not be parsed: %v\n", err)
		return
	}

	vulnFixMetric.Reset()
	vulnMetric.Reset()

	for _, v := range vulnerabilities {
		severity, ok := v["severity"].(string)
		if !ok {
			fmt.Println("Severity value was not read in vulnerabilities.")
			continue
		}

		countWithFix, ok := v["countWithFix"].(float64)
		if !ok {
			fmt.Println("CountWithFix value was not read in vulnerabilities.")
			continue
		}
		countTotal, ok := v["countTotal"].(float64)
		if !ok {
			fmt.Println("countTotal value was not read in vulnerabilities.")
			continue
		}

		vulnFixMetric.WithLabelValues(severity).Set(countWithFix)
		vulnMetric.WithLabelValues(severity).Set(countTotal)
	}
}
