package application

import (
	"encoding/json"
	"fmt"
	"os"
)

func AppVuln(namespaces string) []byte {

	apiEndpoint = os.Getenv("API_ENDPOINT")
	if apiEndpoint == "" {
		fmt.Println("API_ENDPOINT not read env")
		os.Exit(1)
	}
	data, err := fetchData(apiEndpoint + "/api/applications?applicationEnvs[containElements]=" + namespaces + "&page=1&pageSize=50&sortKey=vulnerabilities&sortDir=DESC")
	if err != nil {
		fmt.Printf("Failed to retrieve data from API: %v\n", err)
		return nil
	}

	var jsonData JSONData
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Printf("JSON data could not be parsed: %v\n", err)
		return nil
	}

	for _, app := range jsonData.Items {
		appNamespace := namespaces
		applicationNamespacePackages := app.Packages
		applicationNamespaceName := app.ApplicationName
		applicationNamespaceType := app.ApplicationType
		for _, vulnerability := range app.Vulnerabilities {
			vulnerabilitiesSeverity := vulnerability.Severity
			vulnerabilitiesCount := vulnerability.Count
			applicationVulnerabilities.WithLabelValues(appNamespace, applicationNamespaceName, applicationNamespaceType, vulnerabilitiesSeverity).Set(float64(vulnerabilitiesCount))
		}

		for _, cisBenchmark := range app.CisDockerBenchmarkResults {
			cisLevel := cisBenchmark.Level
			cisCount := cisBenchmark.Count
			applicationCisBenchmark.WithLabelValues(appNamespace, applicationNamespaceName, cisLevel).Set(float64(cisCount))
		}

		applicationVulnerabilitiesPackages.WithLabelValues(appNamespace, applicationNamespaceName).Set(float64(applicationNamespacePackages))
	}

	return nil
}
