package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(test *testing.T) {

	request, err := http.NewRequest("GET", "", nil)
	if err != nil {
		test.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	httpHandler := http.HandlerFunc(handler)

	httpHandler.ServeHTTP(recorder, request)

	// Check status code
	if status := recorder.Code; status != http.StatusOK {
		test.Errorf("Handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	expected := "Hello World!"
	actual := recorder.Body.String()
	if actual != expected {
		test.Errorf("Handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestRouter(test *testing.T) {
	router := newRouter()

	mockserver := httptest.NewServer(router)

	response, err := http.Get(mockserver.URL + "/hello")
	if err != nil {
		test.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		test.Errorf("Status should be ok, got %d", response.StatusCode)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		test.Fatal(err)
	}
	responseString := string(body)
	expected := "Hello World!"

	if responseString != expected {
		test.Errorf("Response should be %s, got %s", expected, responseString)
	}
}

func TestRouterForNonExistentRoute(test *testing.T) {
	router := newRouter()
	mockserver := httptest.NewServer(router)

	// Post
	response, err := http.Post(mockserver.URL+"/hello", "", nil)
	if err != nil {
		test.Fatal(err)
	}

	// http code / msg
	if response.StatusCode != http.StatusMethodNotAllowed {
		test.Errorf("Status should be 405, got %d", response.StatusCode)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		test.Fatal(err)
	}
	responseString := string(body)

	// expect ""
	expected := ""

	if responseString != expected {
		test.Errorf("Response should be %s, got %s", expected, responseString)
	}
}

func TestStaticFileServer(test *testing.T) {
	router := newRouter()
	mockserver := httptest.NewServer(router)

	response, err := http.Get(mockserver.URL + "/assets/")
	if err != nil {
		test.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		test.Errorf("Status should be 200, got %d", response.StatusCode)
	}

	// Do not test the content of the index.html file
	// Only test the content type
	contenType := response.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"
	if expectedContentType != contenType {
		test.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contenType)
	}
}
