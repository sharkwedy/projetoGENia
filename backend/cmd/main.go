package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/subscriptions", CreateSubscription).Methods("POST")
	r.HandleFunc("/api/subscriptions/{id}", GetSubscription).Methods("GET")
	r.HandleFunc("/api/subscriptions/{id}", UpdateSubscription).Methods("PUT")
	r.HandleFunc("/api/subscriptions/{id}", DeleteSubscription).Methods("DELETE")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
