package main

import (
	"time"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", GetData).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetData(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(time.Now())
}
