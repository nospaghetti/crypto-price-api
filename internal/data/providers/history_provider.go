package providers

import "github.com/nospaghetti/crypto-price-api/internal/models"

type HistoryProvider interface {
	GetHistory() []models.Price
}
