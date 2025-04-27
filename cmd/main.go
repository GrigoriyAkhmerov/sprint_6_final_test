package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	server := server.Router(logger)
	if err := server.StartServer(); err != nil {
		server.Logger.Fatalf("starting server error: %s", err)
	}
}
