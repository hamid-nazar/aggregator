package main

import (
	"net/http"
)

func healthCheckHandler(writer http.ResponseWriter, request *http.Request) {

	type statusResponse struct {
		Status string `json:"status"`
	}

	status := statusResponse{Status: "OK"}

	respondWithJSON(writer, 200, status)
}
