package main

import (
	"log"
	"net/http"
	"math/rand"

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

func main() {
	r := prometheus.NewRegistry()
	r.MustRegister(onlineUsers)

	go func() {
		for {
			onlineUsers.Set(float64(rand.Intn(2000)))
		}
	}()

	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8181", nil))
}
