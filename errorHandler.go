package main

import (
	"net/http"
)

func errorHandler(writer http.ResponseWriter, request *http.Request) {

	respondWithError(writer, 400, "Something went wrong. Please try again later.")
}
