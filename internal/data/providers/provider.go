package providers

import "github.com/nospaghetti/crypto-price-api/internal/models"

type Provider interface {
	GetHistory() ([]models.Price, error)
	GetPrices() ([]models.Price, error)
	GetName() string
}
