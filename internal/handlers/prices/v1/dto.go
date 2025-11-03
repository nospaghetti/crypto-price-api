package v1

import "github.com/nospaghetti/crypto-price-api/internal/models"

type GetPricesRequest struct {
	Symbols    []string `json:"symbols"`
	Currencies []string `json:"currencies"`
}

type GetPricesResponse struct {
	Prices    map[string]models.Price `json:"prices"`
	FromCache bool                    `json:"from_cache"`
}
