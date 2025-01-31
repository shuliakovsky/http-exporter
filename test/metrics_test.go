package test

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"http-exporter/internal/metrics"
)

func TestMetricsInit(t *testing.T) {
	metrics.Init()
	if testutil.ToFloat64(metrics.HTTPStatus.WithLabelValues("test_url")) != 0 {
		t.Error("Expected initial metric value to be 0")
	}
}
