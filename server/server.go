package server

import (
	"log"
	"multiplier/handler"
	"net/http"
	"time"
)

type Server struct {
	router *http.ServeMux
}

func NewServer(generatorHandler handler.GeneratorHandler) *Server {
	router := http.NewServeMux()

	generatorHandler.RegisterRoutes(router)

	return &Server{router: router}
}

func (s *Server) Run(port string) {
	srv := &http.Server{
		Handler:      s.router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(srv.ListenAndServe())
}
