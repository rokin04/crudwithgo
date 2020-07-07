package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()
	services(r)
	// Server
	log.Fatal(http.ListenAndServe(":8080", r))
}
