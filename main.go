package main

import (
	"eduSense/controllers"
	"eduSense/helpers"
	// Import the gorilla/mux library we just installed
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	helpers.Route("/students", controllers.StudentController{}, r)

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return r
}

func main() {
	r := newRouter()
	fmt.Println("Start listening server 8080")
	http.ListenAndServe(":8080", r)
}

