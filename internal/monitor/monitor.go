package monitor

import (
	"net/http"
	"time"

	"http-exporter/internal/metrics"
)

func MonitorURL(url string, interval time.Duration) {
	for {
		resp, err := http.Get(url)
		if err != nil {
			metrics.HTTPStatus.WithLabelValues(url).Set(0)
		} else {
			metrics.HTTPStatus.WithLabelValues(url).Set(float64(resp.StatusCode))
			resp.Body.Close()
		}
		time.Sleep(interval)
	}
}
