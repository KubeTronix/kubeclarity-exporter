package metrics

import (
	"kubetronix/probe/application"
	counter "kubetronix/probe/counter"
)

func GetMetrics() {
	counter.GetCounters()
	counter.VulnCounters()
	counter.LicenseCounters()
	counter.LangCounters()
	application.NamespaceCounters()
}
