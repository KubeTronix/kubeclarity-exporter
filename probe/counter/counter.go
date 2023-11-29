package counter

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetCounters() {

	apiEndpoint = os.Getenv("API_ENDPOINT")
	if apiEndpoint == "" {
		fmt.Println("API_ENDPOINT not read env")
		os.Exit(1)
	}
	data, err := fetchData(apiEndpoint + "/api/dashboard/counters")
	if err != nil {
		fmt.Printf("Failed to retrieve data from API: %v\n", err)
		return
	}

	var counters map[string]int
	err = json.Unmarshal(data, &counters)
	if err != nil {
		fmt.Printf("JSON data could not be parsed: %v\n", err)
		return
	}

	applicationsMetric.Reset()
	applicationsMetric.WithLabelValues().Set(float64(counters["applications"]))

	packagesMetric.Reset()
	packagesMetric.WithLabelValues().Set(float64(counters["packages"]))

	resourcesMetric.Reset()
	resourcesMetric.WithLabelValues().Set(float64(counters["resources"]))
}
