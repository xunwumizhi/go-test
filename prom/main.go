package main

import (
	"math"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	myCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sleep_total",
		Help: "The total number of sleep times",
	})

	temps = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "pond_temperature_celsius",
			Help:       "The temperature of the frog pond.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"species"},
	)
)

func collector() {
	go func() {
		for {
			myCounter.Inc()
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		// Simulate some observations.
		for i := 0; i < 1000; i++ {
			temps.WithLabelValues("litoria-caerulea").Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)
			temps.WithLabelValues("lithobates-catesbeianus").Observe(32 + math.Floor(100*math.Cos(float64(i)*0.11))/10)
		}
		temps.WithLabelValues("only-one").Observe(233)
	}()

}

func init() {
	// Create a Summary without any observations.
	temps.WithLabelValues("no-observations")
}

func main() {
	collector()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
