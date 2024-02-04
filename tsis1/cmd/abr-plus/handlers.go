package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	abrplus "github.com/tttulepbergen/tsis1/pkg/abr-plus"
)

type Response struct {
	Movies []abrplus.Movie `json:"movies"`
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok", "info": abrplus.Info()})
}

func movies(w http.ResponseWriter, r *http.Request) {
	movies := abrplus.GetMovies()
	respondWithJSON(w, http.StatusOK, movies)
}

func movie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	movie, err := abrplus.GetMovie(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	respondWithJSON(w, http.StatusOK, movie)
}
