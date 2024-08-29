package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respondWithJSON(writer http.ResponseWriter, code int, payload interface{}) {

	response, err := json.Marshal(payload)
	fmt.Printf("Response: %s\n", response)
	if err != nil {
		fmt.Printf("Failed to convert payload to JSON. Err: %s", err)

		writer.WriteHeader(500)
		return
	}
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(response)
}

func respondWithError(writer http.ResponseWriter, code int, message string) {

	if code > 499 {
		log.Printf("Responding with 5XX error: %s", message)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(writer, code, errorResponse{Error: message})
}
