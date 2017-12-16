package collector

import (
        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/common/log"
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
        log.Infof("scrape target: %s", c.Target)
        start := time.Now()
        time.Sleep(time.Second)
        duration := time.Since(start).Second()
        ch <- prometheus.MustNewConstMetric(scrapeDuraiton, prometheus.GaugeValue, duration)
}
