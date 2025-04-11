package main

import (
	"log"
	"net/http"
	"os"

	"ia-go-comment-fetcher/db"
	"ia-go-comment-fetcher/routes"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	log.Println(".env loaded successfully")
	log.Println("MONGO_URI =", os.Getenv("MONGO_URI"))
}

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	mongoClient := db.InitMongo(mongoURI)

	routes.RegisterRoutes(mongoClient)

	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}