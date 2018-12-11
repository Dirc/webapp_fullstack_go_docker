package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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
	// Enable/disable the use of an external database by comment in/out the lines UNTIL HERE.

	// Connect to: postgres container
	connString := "host=db user=postgres password=secret dbname=bird_encyclopedia sslmode=disable"
	// Connect to: local Postgres database
	//connString := "dbname=localhost sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	InitStore(&dbStore{db: db})

	// UNTIL HERE

	router := newRouter()
	http.ListenAndServe(":8080", router)
}
