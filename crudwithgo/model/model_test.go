package model

import (
	"encoding/json"
	// "fmt"

	// "io"
	// "io"
	"io/ioutil"
	// "net/http"
	// "net/http"
	"net/http/httptest"
	"testing"
	// "time"

	// "github.com/gorilla/mux"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// func TestGetUser(t *testing.T) {
// 	router := mux.NewRouter() //initialise the router
// 	testServer := httptest.NewServer(router) //setup the testing server
// 	res, err := http.Get(fmt.Sprintf("http://localhost:8080/api/users/1"))
// 	body, _ := ioutil.ReadAll(res.Body)
// 	fmt.Println(res.StatusCode, testServer.URL)
// 	if err != nil {
// 		t.Fatalf("could not send GET req")
// 	} else {

// 	}
// 	if res.StatusCode == 200 {
// 		fmt.Println("got resp", string(body))
// 	} else {
// 		t.Fatalf("got status error")
// 	}
// 	type Mock struct {
// 		ID int  `json:id`
// 		Name string  `json:name`
// 		Birthday string  `json:birthday`
// 		Onboarded string  `json:onboarded`
// 		IsActive bool  `json:isactive`
// 	}
// 	data := []Mock{{ ID: 1, Name: "Karthik", Birthday: time.Now().Format("01-01-2006"), Onboarded: time.Now().Format("01-01-2005"), IsActive: true}}
// 	assert.Equal(t, data, string(body))
// }

func TestGetuser(t *testing.T) {
	type Mock struct {
		ID int  `json:id`
		Name string  `json:name`
		Birthday string  `json:birthday`
		Onboarded string  `json:onboarded`
		IsActive bool  `json:isactive`
	}

	params := map[string]string {
		"id": "1",
	}

	newReq := httptest.NewRequest("GET", "http://localhost:8080/api/users", nil)
	req := mux.SetURLVars(newReq, params)
	w := httptest.NewRecorder()
	GetUser(w, req);

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	userMock := Mock{}
	json.Unmarshal(body, &userMock)

	if 200 != resp.StatusCode {
		t.Fatal("status code not good")
	}
	assert.Equal(t, 1 , userMock.ID)
}