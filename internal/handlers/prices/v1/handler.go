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

// GetPrices     GetPrice godoc
// @Summary      Current crypto price
// @Description  Returns latest crypto price for given symbol and currency
// @Tags         prices
// @Param        symbol   path      string  true  "Crypto symbol (e.g. btc, eth)"
// @Param        currency query     string  false "Fiat currency (default: usd)"
// @Router       /prices/{symbol} [get]
func (h *PricesHandler) GetPrices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
