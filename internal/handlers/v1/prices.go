package v1

import (
	"encoding/json"
	"net/http"
	"strings"

	_ "github.com/nospaghetti/crypto-price-api/internal/dto/v1"
	"github.com/nospaghetti/crypto-price-api/internal/services"
)

type PricesHandler struct {
	service *services.PricesService
}

func NewPricesHandler(service *services.PricesService) *PricesHandler {
	return &PricesHandler{service}
}

type APIError struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func writeAPIError(w http.ResponseWriter, status int, code, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	var e APIError
	e.Error.Code = code
	e.Error.Message = msg
	_ = json.NewEncoder(w).Encode(e)
}

// GetPrices GetPrice godoc
// @Summary Current crypto prices
// @Description Returns latest crypto prices for given symbols and currencies. Only found fiat currencies are returned.
// @Tags prices
// @Param symbol query string true "Crypto symbols (e.g. btc, eth)"
// @Param currency query string false "Fiat currencies (default: usd)"
// @Success 200 {object} v1.GetPricesDTO
// @Router /prices [get]
func (h *PricesHandler) GetPrices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		symbol := strings.TrimSpace(r.URL.Query().Get("symbol"))
		if symbol == "" {
			writeAPIError(w, http.StatusBadRequest, "invalid_argument", "symbol query parameter is required")
			return
		}

		fiatsParam := r.URL.Query().Get("fiats")
		if fiatsParam == "" {
			http.Error(w, "fiats query parameter is required", http.StatusBadRequest)
			return
		}
		rawFiats := strings.Split(fiatsParam, ",")
		var fiats []string

		for _, fiat := range rawFiats {
			if fiat = strings.TrimSpace(fiat); fiat != "" {
				fiats = append(fiats, fiat)
			}
		}

		prices, err := h.service.GetPrices(symbol, fiats)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(prices); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
