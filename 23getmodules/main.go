package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greet()

	r := mux.NewRouter()

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/about", serveAbout).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greet() {
	fmt.Println("Server GO")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	w.Write([]byte("<body style=\"background-color: #212121; color: white\"><h1 style=\"text-align: center\">Home Route</h1></body>"))
}

func serveAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About")
	w.Write([]byte("<body style=\"background-color: #212121; color: white\"><h1 style=\"text-align: center\">About Route</h1></body>"))
}
