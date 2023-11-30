package application

import (
	"encoding/json"
	"fmt"
	"os"
)

func NamespaceCounters() {

	apiEndpoint = os.Getenv("API_ENDPOINT")
	if apiEndpoint == "" {
		fmt.Println("API_ENDPOINT not read env")
		os.Exit(1)
	}
	data, err := fetchData(apiEndpoint + "/api/namespaces")
	if err != nil {
		fmt.Printf("Failed to retrieve data from API: %v\n", err)
		return
	}

	var namespaces []map[string]interface{}
	err = json.Unmarshal(data, &namespaces)
	if err != nil {
		fmt.Printf("JSON data could not be parsed: %v\n", err)
		return
	}

	for _, v := range namespaces {
		namespaces, ok := v["name"].(string)

		if !ok {
			fmt.Println("Count value was not read in namespaces.")
			continue
		}
		AppVuln(namespaces)
	}
}
