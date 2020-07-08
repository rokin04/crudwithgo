package main

import (
	"log"
	"net/http"
	"github.com/rokin04/crudwithgo/crudwithgo/services"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	// Init router
	r := mux.NewRouter()
	services.Services(r)

	//Logiing request and response status

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(r)


	// Server
	log.Fatal(http.ListenAndServe(":8080", n))
}
