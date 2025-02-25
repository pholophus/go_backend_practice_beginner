package main

import (
	"log"
	"net/http"

	"github.com/pholophus/go_backend_practice_beginner/internal/config"
	"github.com/pholophus/go_backend_practice_beginner/internal/routes"
)

func main() {
	cfg := config.GetConfig()

	router := routes.SetupRouter()

	log.Printf("Server is running on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}