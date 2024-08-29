package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT is not found in .env file")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Use(middleware.Logger)

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", healthCheckHandler)
	v1Router.Get("/error", errorHandler)

	router.Mount("/api/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", port),
	}

	fmt.Printf("Server is running on port %v", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
