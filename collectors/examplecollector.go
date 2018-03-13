package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"runtime/debug"
	"time"
)

const namespace = "example"

var (
	scrapeDuration = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "subsystem", "scrape_duration_seconds"),
		"Duration of a collector scrape",
		nil,
		nil,
	)
)

type ExampleCollector struct {
	Target string
}

func (c ExampleCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- scrapeDuration
}

func (c ExampleCollector) Collect(ch chan<- prometheus.Metric) {
	defer func() {
		if e, ok := recover().(error); ok {
			log.Errorf("WARN: panic in %v", e)
			debug.PrintStack()
		}
	}()
	log.Debugf("scrape target: %s", c.Target)
	start := time.Now()
	time.Sleep(time.Second)
	duration := time.Since(start).Seconds()
	ch <- prometheus.MustNewConstMetric(scrapeDuration, prometheus.GaugeValue, duration)
}
