package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, status_code int, msg string) {
	if status_code > 499 {
		log.Println("Internal server error: ", msg)
	}
	type errReponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, status_code, errReponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, status_code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", payload) // log payload instead of error
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status_code)
	w.Write(dat)
}
