package main

import (
	"exporter-template/collectors"
	"exporter-template/collectors/api"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
	_ "net/http/pprof"
)

func runCollector(collector prometheus.Collector, w http.ResponseWriter, r *http.Request) {
	registry := prometheus.NewRegistry()
	registry.MustRegister(collector)
	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var collectortype prometheus.Collector
	variables := mux.Vars(r)
	ctype := variables["type"]
	target := r.URL.Query().Get("target")
	if target == "" {
		http.Error(w, "'target' must be specified", http.StatusBadRequest)
		return
	}
	switch ctype {
	case "type1":
		collectortype = collector.ExampleCollector{Target: target}
	default:
		http.Error(w, "unsupported metric path", http.StatusNotFound)
		return
	}
	runCollector(collectortype, w, r)
}

func main() {
	var (
		listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry").Default(":9100").String()
	)
	kingpin.Parse()
	r := mux.NewRouter()
	r.PathPrefix("/debug/pprof").Handler(http.DefaultServeMux)
	r.HandleFunc("/api/v1/resource", api.GetResource)
	r.HandleFunc("/{type}", handler)
	log.Infof("Listening on %s", *listenAddress)
	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
