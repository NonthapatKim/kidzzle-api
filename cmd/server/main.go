// @title kidzzle API
// @version 1.0
// @description This is the API documentation for the kidzzle project.
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
package main

import (
	"log"

	"github.com/NonthapatKim/kidzzle-api/internal/server"
)

func main() {
	if err := server.RunServer(); err != nil {
		log.Fatal(err)
	}
}
