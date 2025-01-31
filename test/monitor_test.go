package test

import (
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"http-exporter/internal/metrics"
	"http-exporter/internal/monitor"
)

func TestMonitorURL(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "http://example.com"
	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, "OK"))

	interval := 1 * time.Second

	go monitor.MonitorURL(url, interval)
	time.Sleep(2 * time.Second)

	metric := testutil.ToFloat64(metrics.HTTPStatus.WithLabelValues(url))
	if metric != 200 {
		t.Errorf("Expected HTTP status 200, got %v", metric)
	}
}
