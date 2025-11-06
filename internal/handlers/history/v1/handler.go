package v1

import (
	"net/http"

	"github.com/nospaghetti/crypto-price-api/internal/services"
)

type HistoryHandler struct {
	service *services.HistoryService
}

func NewHistoryHandler(service *services.HistoryService) *HistoryHandler {
	return &HistoryHandler{service: service}
}

func (h *HistoryHandler) GetHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
