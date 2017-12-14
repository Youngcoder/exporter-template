package main

import (
        "fmt"
        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promhttp"
        "github.com/prometheus/common/log"
        "gopkg.in/alecthomas/kingpin.v2"
        "net/http"
        _ "net/http/pprof"
        "strings"
        "exporter-template/config"
        "exporter-template/collectors"
)

func init() {
        config.GetDBHandle()
}

func runCollector(collector prometheus.Collector, w http.ResponseWrite, r *http.Request) {
        registry := prometheus.NewRegistry()
        registry.MustRegister(collector)
        h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
        h.ServeHTTP(w, r)
}

func handler(w http.ResponseWrite, r *http.Request) {
        var collectortype prometheus.Collector
        target := r.URL.Query().Get("target")
        if target == "" {
                http.Error(w, "missing target", 400)
                return
        }
        switch strings.Split(fmt.Sprintf("%s", r.URL), "?")[0] {
        case "/metricspath":
                collectortype = collector.ExampleCollector{Target: target}
        }
        runCollector(collectortype, w, r)
}

func main() {
        var (
                listenAddress = kingpin.Flag("web.listen-address","Address to listen on for web interface and telemetry").Default(":9100").String()
        )
        kingpin.Parse()
        http.HandleFunc("/metricspath", handler)
        log.Infof("Listening on %s", *listenAddress)
        err := http.ListenAndServe(*listenAddress, nil)
        if err != nil {
                log.Fatal(err)
        }
        defer config.CloseDBHandle()
}
