package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/subscriptions", CreateSubscription).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Subscription created!"))
}