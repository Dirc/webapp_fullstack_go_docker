package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBirdHandlertest(test *testing.T) {

	birds = []Bird{
		{"kanarie", "Yellow bird"},
	}

	request, err := http.NewRequest("GET", "", nil)
	if err != nil {
		test.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(getBirdHandler)

	hf.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := Bird{"kanarie", "Yellow bird"}
	b := []Bird{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		test.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		test.Errorf("Response should be %v, got %v", expected, actual)
	}

}
