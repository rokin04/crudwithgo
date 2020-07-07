package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

// User struct
type User struct {
	ID int  `json:id`
	Name string  `json:name`
	Birthday string  `json:birthday`
	Onboarded string  `json:onboarded`
	IsActive bool  `json:isactive`
}

// Mock data
var users []User

// GET all users

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Get user By ID

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // set params
	i, _ := strconv.Atoi(params["id"])
	fmt.Println(i)
	for _, item := range users {
		fmt.Println(item.ID)
		if item.ID == i {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

// Create User

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = rand.Intn(10000) // mock ID
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// Update User

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range users {
		if item.ID == id {
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = id
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
		}
	}
	json.NewEncoder(w).Encode(users)
}

// Delete User

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range users {
		if item.ID == id {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {

	// Init router
	
	r := mux.NewRouter()

	// Mock data
	users = append(users, User{ID: 1, Name: "Karthik", Birthday: time.Now().Format("01-01-2006"), Onboarded: time.Now().Format("01-01-2005"), IsActive: true})
	users = append(users, User{ID: 2, Name: "Rokin", Birthday: time.Now().Format("01-01-2012"), Onboarded: time.Now().Format("01-01-2015"), IsActive: false})


	// Endpoints

	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	// Server

	log.Fatal(http.ListenAndServe(":8080", r))
}	