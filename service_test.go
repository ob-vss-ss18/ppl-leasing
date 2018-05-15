package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func Test(t *testing.T) {
	//Generate Request
	req, err := http.NewRequest("GET", "/query", nil);
	if err != nil {
		t.Fatal(err);
	}

	//ResponseRecorder
	recorder := httptest.NewRecorder();
	handler := http.HandlerFunc(sayhelloName);

	handler.ServeHTTP(recorder, req);

	//Check status
	if recorder.Code != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", recorder.Code, http.StatusOK);
	}

	//Check body
	expected := "Hello world, Dev Heroku!";
	if recorder.Body.String() != expected {
		t.Errorf("returned false body: got %v want %v", recorder.Body.String(), expected);
	}

}
