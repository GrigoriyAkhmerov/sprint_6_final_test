package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Point 1. Creating server structure.
type Server struct {
	Logger     *log.Logger
	HttpServer *http.Server
}

// Point 2. Router function creating http-router.
func (s *Server) StartServer() error {
	s.Logger.Println("Server is running")
	return s.HttpServer.ListenAndServe()
}

// Point 2. Router function creating http-router.
func Router(logger *log.Logger) *Server {

	// Point 3. Register handlers.
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", handlers.UploadEndpointHandler)
	mux.HandleFunc("/", handlers.RootEndpointHandler)

	// Point 4. Creating http.Server structure exemplar.
	exemplarOfSever := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     logger,
		HttpServer: exemplarOfSever,
	}
}
