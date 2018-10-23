package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()

	// URI: "/hello" return handler func
	router.HandleFunc("/hello", handler).Methods("GET")

	// URI: "/assets/" return index.html
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets", http.FileServer(staticFileDirectory))
	router.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	// URI: "/bird" enable getBirdHandler and createBirdHandler (from bird_handlers.go)
	router.HandleFunc("/bird", getBirdHandler).Methods("GET")
	router.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return router
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	router := newRouter()
	http.ListenAndServe(":8080", router)
}
