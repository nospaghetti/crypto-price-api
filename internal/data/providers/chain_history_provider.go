package providers

import "github.com/nospaghetti/crypto-price-api/internal/models"

type ChainHistoryProvider struct {
	providers []HistoryProvider
}

func NewChainHistoryProvider(providers []HistoryProvider) *ChainHistoryProvider {
	return &ChainHistoryProvider{providers}
}

func (c *ChainHistoryProvider) GetHistory() []models.Price {
	for _, provider := range c.providers {
		result := provider.GetHistory()

		if len(result) > 0 {
			return result
		}
	}

	return nil
}
