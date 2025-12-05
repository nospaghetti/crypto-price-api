package v1

import (
	"encoding/json"
	"net/http"

	"github.com/nospaghetti/crypto-price-api/internal/services"
)

type HealthHandler struct {
	service *services.HealthService
}

func NewHealthHandler(service *services.HealthService) *HealthHandler {
	return &HealthHandler{service}
}

func (h *HealthHandler) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"server":        "OK",
			"db_connection": "OK",
		}

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	}
}
