package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":8080"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health-check", healthCheck)
	r.HandleFunc("/movies", movies)
	r.HandleFunc("/movies/{id:[0-9]+}", movie)

	log.Printf("Starting server on %s\n", PORT)
	err := http.ListenAndServe(PORT, r)
	log.Fatal(err)
}
