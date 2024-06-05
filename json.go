package main

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, status_code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {

	}
}
