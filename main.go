package main

import (
	"log"
	"net/http"

	"ia-go-comment-fetcher/routes"
)

func main() {
	routes.RegisterRoutes()

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}