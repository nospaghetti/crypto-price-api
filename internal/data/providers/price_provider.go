package providers

import "github.com/nospaghetti/crypto-price-api/internal/models"

type PriceProvider interface {
	GetPrices() []models.Price
}
