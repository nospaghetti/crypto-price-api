package providers

import "github.com/nospaghetti/crypto-price-api/internal/models"

type ChainPriceProvider struct {
	providers []PriceProvider
}

func NewChainPriceProvider(providers []PriceProvider) *ChainPriceProvider {
	return &ChainPriceProvider{providers}
}

func (c *ChainPriceProvider) GetPrices() []models.Price {
	for _, provider := range c.providers {
		result := provider.GetPrices()

		if len(result) > 0 {
			return result
		}
	}

	return nil
}
