package v1

import (
	"net/http"

	"github.com/nospaghetti/crypto-price-api/internal/services"
)

type PricesHandler struct {
	service *services.PricesService
}

func NewPricesHandler(service *services.PricesService) *PricesHandler {
	return &PricesHandler{service}
}

// GetPrices godoc
// @Summary Get current price
// @Description Get current price
// @Tags prices
// @Accept JSON
// @Produce JSON
// @Success 200 {object} models.Price
// @Router /api/v1/prices [get]
func (h *PricesHandler) GetPrices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
