package providers

import "github.com/nospaghetti/crypto-price-api/internal/models"

type CoinGecko struct {
}

func NewCoinGecko() *CoinGecko {
	return &CoinGecko{}
}

func (c *CoinGecko) GetPrices() []models.Price {
	return []models.Price{}
}
