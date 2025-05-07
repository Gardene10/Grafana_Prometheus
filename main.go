package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var onlineUsers = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "goapp_online_users",
	Help: "Online users",
	ConstLabels: map[string]string{
		"website": "ecommerce",
	},
})

var httpRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "goapp_http_requests_total",
		Help: "Total HTTP requests",
	},
	[]string{"method", "path"},
),

var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "goapp_http_request_duration",
	Help: "Duration in seconds of all HTTP requests",
}, []string{"handler"})

func main() {
	r := prometheus.NewRegistry()
	r.MustRegister(onlineUsers)
	r.MustRegister(httpRequests)
	r.MustRegister(httpDuration)

	go func() {
		for {
			onlineUsers.Set(float64(rand.Intn(2000)))
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// Incrementa a métrica com as labels corretas
		httpRequests.WithLabelValues(req.Method, req.URL.Path).Inc()

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8181", nil))
}
