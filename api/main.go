package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "go_request_operations_total",
	Help: "The total number of processed requests",
})

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", GetData).Methods("GET")
	router.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetData(w http.ResponseWriter, r *http.Request) {
	counter.Inc()

	json.NewEncoder(w).Encode(time.Now())
}
