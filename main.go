package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type GeneratorService struct {
	rtp float64
}

func NewGenerator(rtp float64) *GeneratorService {
	return &GeneratorService{rtp: rtp}
}

func (g *GeneratorService) Get() float64 {
	rtp := g.rtp

	u := max(rand.Float64(), 0.000001)

	multiplier := rtp / u
	multiplier = min(10000, max(1, multiplier))

	return multiplier
}

type GeneratorHandler struct {
	generatorService GeneratorService
}

func NewGeneratorHandler(generatorService GeneratorService) *GeneratorHandler {
	return &GeneratorHandler{generatorService: generatorService}
}

func (h *GeneratorHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/get", h.get)
}

func (h *GeneratorHandler) get(w http.ResponseWriter, r *http.Request) {
	response := GeneratorResponse{
		Result: h.generatorService.Get(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}

type GeneratorResponse struct {
	Result float64 `json:"result"`
}

type Server struct {
	router *http.ServeMux
}

func NewServer(generatorHandler GeneratorHandler) *Server {
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

func main() {
	rtp := flag.Float64("rtp", 1, "RTP value")
	flag.Parse()

	// check for rtp flag range
	if *rtp <= 0 || *rtp > 1 {
		log.Fatal("RTP must be in range (0;1.0]\n")
	}

	generatorService := NewGenerator(*rtp)
	generatorHandler := NewGeneratorHandler(*generatorService)

	server := NewServer(*generatorHandler)

	// make it possible to use custom port via env variables
	// it would be more preferable to add a .env file, but for now this solution will do
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Println("SERVER_PORT environment vartiable is not set. Using default port 64333")
		port = "64333"
	}

	server.Run(port)
}
