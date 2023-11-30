package application

import (
	"io/ioutil"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Application struct {
	ApplicationName           string `json:"applicationName"`
	ApplicationResources      int    `json:"applicationResources"`
	ApplicationType           string `json:"applicationType"`
	CisDockerBenchmarkResults []struct {
		Count int    `json:"count"`
		Level string `json:"level"`
	} `json:"cisDockerBenchmarkResults"`
	Environments    []string `json:"environments"`
	ID              string   `json:"id"`
	Labels          []string `json:"labels"`
	Packages        int      `json:"packages"`
	Vulnerabilities []struct {
		Count    int    `json:"count"`
		Severity string `json:"severity"`
	} `json:"vulnerabilities"`
}

type JSONData struct {
	Items []Application `json:"items"`
	Total int           `json:"total"`
}

var (
	apiEndpoint string

	applicationVulnerabilities = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kc_app_vuln",
			Help: "Total Count of Kubernetes application vulnerability in Namespaces",
		},
		[]string{"namespace", "applicationName", "applicationType", "vulnerabilitiesSeverity"},
	)

	applicationVulnerabilitiesPackages = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kc_app_packages",
			Help: "Total Count of Kubernetes application packages in Namespaces",
		},
		[]string{"namespace", "applicationName"},
	)

	applicationCisBenchmark = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kc_app_cis",
			Help: "Total Count of Kubernetes application cisbenchmark in Namespaces",
		},
		[]string{"namespace", "applicationName", "level"},
	)
)

func fetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
