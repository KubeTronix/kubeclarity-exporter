package metrics

import (
	counter "kubetronix/probe/counter"
)

func GetMetrics() {
	counter.GetCounters()
	counter.VulnCounters()
	counter.LicenseCounters()
	counter.LangCounters()
}
