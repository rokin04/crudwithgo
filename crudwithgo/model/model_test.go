package model

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetusers(t *testing.T) {
	handler := func(w http.ResponseWriter, r http.Request) {
		io.WriteString(w, "{\"status\" : \"good\"}")
	}

	req := httptest.NewRequest("GET", "http://localhost:8080/api/users", nil)
	w := httptest.NewRecorder()
	handler(w, *req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	if 200 != resp.StatusCode {
		t.Fatal("status code not good")
	}
}