package handler

import (
	"encoding/json"
	"log"
	"multiplier/service"
	"net/http"
)

type GeneratorHandler struct {
	generatorService service.GeneratorService
}

func NewGeneratorHandler(generatorService service.GeneratorService) *GeneratorHandler {
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
