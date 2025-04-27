package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	server := server.Router(logger)
	err := server.StartServer()
	err != nil {
		server.Logger.Fatalf("Error starting server: %s", err)
	}
}
