package model

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {

	req := httptest.NewRequest("GET", "http://localhost:8080/api/users", nil)
	w := httptest.NewRecorder()
	GetUsers(w, req);

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var userMock []User
	json.Unmarshal(body, &userMock)

	if 200 != resp.StatusCode {
		t.Fatal("status code not good")
	}
	if len(userMock) == 0 {
		t.Fatal("No users found")
	}
}

func TestGetUser(t *testing.T) {

	params := map[string]string {
		"id": "1",
	}

	newReq := httptest.NewRequest("GET", "http://localhost:8080/api/users", nil)
	req := mux.SetURLVars(newReq, params)
	w := httptest.NewRecorder()
	GetUser(w, req);

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var userMock User
	json.Unmarshal(body, &userMock)

	if 200 != resp.StatusCode {
		t.Fatal("status code not good")
	}
	assert.Equal(t, 1 , userMock.ID)
}

func TestCreateUser(t *testing.T) {

	payload := []byte(
		`{ 
			"Name": "newUser", 
			"Birthday": "01-01-2006", 
			"Onboarded": "01-01-2005", 
			"IsActive": true 
		}`,
	)
	newReq := httptest.NewRequest("POST", "http://localhost:8080/api/users", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	CreateUser(w, newReq);

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var userMock User
	json.Unmarshal(body, &userMock)

	if 200 != resp.StatusCode {
		t.Fatal("status code not good")
	}
	assert.Equal(t, "newUser", userMock.Name)
}

func TestUpdateUser(t *testing.T) {

	payload := []byte(
		`{ 
			"Name": "updatedUser", 
			"Birthday": "01-01-2006", 
			"Onboarded": "01-01-2005", 
			"IsActive": true 
		}`,
	)

	params := map[string]string {
		"id": "1",
	}

	newReq := httptest.NewRequest("PUT", "http://localhost:8080/api/users", bytes.NewBuffer(payload))
	req := mux.SetURLVars(newReq, params)
	w := httptest.NewRecorder()
	UpdateUser(w, req);

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var userMock []User
	json.Unmarshal(body, &userMock)
	if 200 != resp.StatusCode {
		t.Fatal("status code not good")
	}
	for _, user := range userMock {
		if user.ID == 1 {
			assert.Equal(t, "updatedUser" , user.Name)
		}
	}
}

func TestDeleteUser(t *testing.T) {

	params := map[string]string {
		"id": "1",
	}

	newReq := httptest.NewRequest("DELETE", "http://localhost:8080/api/users", nil)
	req := mux.SetURLVars(newReq, params)
	w := httptest.NewRecorder()
	DeleteUser(w, req);

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var userMock []User
	json.Unmarshal(body, &userMock)
	if 200 != resp.StatusCode {
		t.Fatal("status code not good")
	}
	for _, user := range userMock {
		if user.ID == 1 {
			t.Fatal("User still exists")
			break
		}
	}
	assert.True(t, true)
}
