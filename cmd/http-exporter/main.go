package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"http-exporter/internal/config"
	"http-exporter/internal/metrics"
	"http-exporter/internal/monitor"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Version and CommitHash will be set during the build process
var Version string = "0.0.1" // Default Version value
var CommitHash string = ""

func printVersion() {
	fmt.Printf("http-exporter version: %s\n", Version)
	if CommitHash != "" {
		fmt.Printf("commit hash: %s\n", CommitHash)
	}
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  --config <path to config file>     Specify the path to the configuration file.")
	fmt.Println("  --help                             Display this help message.")
	fmt.Println("  --version                          Display the current version of the application.")
}

func main() {
	configFile := flag.String("config", "config.yaml", "Path to the configuration file")
	versionFlag := flag.Bool("version", false, "Print the version and exit")
	helpFlag := flag.Bool("help", false, "Print help message")

	flag.Parse()

	if *helpFlag {
		printHelp()
		return
	}

	if *versionFlag {
		printVersion()
		return
	}

	metrics.Init() // Initialize metrics

	cfg, err := config.ReadConfig(*configFile)
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	interval, err := time.ParseDuration(cfg.Interval)
	if err != nil {
		log.Fatalf("error parsing interval: %v", err)
	}

	for _, url := range cfg.URLs {
		go monitor.MonitorURL(url, interval)
	}

	http.Handle("/metrics", promhttp.Handler())
	addr := fmt.Sprintf("%s:%d", cfg.Interface, cfg.Port)
	log.Printf("Starting server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
