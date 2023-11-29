package counter

import (
	"encoding/json"
	"fmt"
	"os"
)

func LangCounters() {

	apiEndpoint = os.Getenv("API_ENDPOINT")
	if apiEndpoint == "" {
		fmt.Println("API_ENDPOINT not read env")
		os.Exit(1)
	}
	data, err := fetchData(apiEndpoint + "/api/dashboard/packagesPerLanguage")
	if err != nil {
		fmt.Printf("Failed to retrieve data from API: %v\n", err)
		return
	}

	var packagesPerLanguage []map[string]interface{}
	err = json.Unmarshal(data, &packagesPerLanguage)
	if err != nil {
		fmt.Printf("JSON data could not be parsed: %v\n", err)
		return
	}

	langMetric.Reset()

	for _, v := range packagesPerLanguage {
		language, ok := v["language"].(string)
		if !ok {
			fmt.Println("Language value was not read in packagesPerLanguage.")
			continue
		}

		perLanguage, ok := v["count"].(float64)
		if !ok {
			fmt.Println("Count value was not read in packagesPerLanguage.")
			continue
		}

		langMetric.WithLabelValues(language).Set(perLanguage)
	}
}
