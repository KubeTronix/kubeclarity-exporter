package counter

import (
	"io/ioutil"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	apiEndpoint string

	vulnMetric = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kc_counter_vuln",
			Help: "Total Count of Vulnerabilities",
		},
		[]string{"severity"},
	)

	vulnFixMetric = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kc_counter_vuln_fix",
			Help: "Total Count of Vulnerabilities",
		},
		[]string{"severity"},
	)

	applicationsMetric = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kc_counter_applications",
			Help: "Toplam uygulama sayısı",
		},
		[]string{},
	)

	packagesMetric = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kc_counter_packages",
			Help: "Toplam paket sayısı",
		},
		[]string{},
	)

	resourcesMetric = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kc_counter_resources",
			Help: "Toplam kaynak sayısı",
		},
		[]string{},
	)

	licenseMetric = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kc_counter_license",
			Help: "Total Count of Vulnerabilities",
		},
		[]string{"license"},
	)

	langMetric = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kc_counter_language",
			Help: "Total Count of Vulnerabilities",
		},
		[]string{"language"},
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
