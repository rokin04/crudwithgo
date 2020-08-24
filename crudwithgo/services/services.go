package services

import (
	"github.com/rokin04/crudwithgo/crudwithgo/model"
	"github.com/gorilla/mux"
)

func Services(r *mux.Router) *mux.Router {
	
	// Endpoints

	r.HandleFunc("/api/users", model.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", model.GetUser).Methods("GET")
	r.HandleFunc("/api/users", model.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", model.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", model.DeleteUser).Methods("DELETE")

	return r
}