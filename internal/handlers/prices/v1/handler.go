package v1

import "net/http"

type HandlerV1 struct {
}

// GetPrices godoc
// @Summary Get current price
// @Description Get current price
// @Tags prices
// @Accept JSON
// @Produce JSON
// @Success 200 {object} models.Price
// @Router /api/v1/prices [get]
func (h *HandlerV1) GetPrices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
