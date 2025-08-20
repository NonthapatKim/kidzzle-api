// @title kidzzle API
// @version 1.0
// @description This is the API documentation for the kidzzle project.
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
package main

import (
	"log"
	"os"

	"github.com/NonthapatKim/kidzzle-api/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file not loaded")
		}
	}

	if err := server.RunServer(); err != nil {
		log.Fatal(err)
	}
}
