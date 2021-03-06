package model

import (
	"encoding/json"
	"fmt"
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
var users = []User{
	User { ID: 1, Name: "Karthik", Birthday: time.Now().Format("01-01-2006"), Onboarded: time.Now().Format("01-01-2005"), IsActive: true },
	User {ID: 2, Name: "Rokin1", Birthday: time.Now().Format("01-01-2012"), Onboarded: time.Now().Format("01-01-2015"), IsActive: false},
}



// GetUsers details
	func GetUsers(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}

// GetUser By ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // set params
	i, _ := strconv.Atoi(params["id"])
	for _, item := range users {
		fmt.Println(item.ID)
		if item.ID == i {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

// CreateUser adds new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = rand.Intn(10000) // mock ID
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// UpdateUser updates existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var user User
	for index, item := range users {
		if item.ID == id {
			users = append(users[:index], users[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = id
			users = append(users, user)
		}
	}
	json.NewEncoder(w).Encode(users)
}

// DeleteUser deletes user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
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