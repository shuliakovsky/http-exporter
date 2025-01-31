package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	HTTPStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_status",
			Help: "HTTP status codes for monitored URLs",
		},
		[]string{"url"},
	)
)

func Init() {
	prometheus.MustRegister(HTTPStatus)
}
