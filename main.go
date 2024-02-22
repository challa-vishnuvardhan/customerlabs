package main

import (
	"log"
	"net/http"

	"customerlabs/app"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", app.Worker).Methods("POST")
	log.Fatal(http.ListenAndServe(":7000", r))
}
