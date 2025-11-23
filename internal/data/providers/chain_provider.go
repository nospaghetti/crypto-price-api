package providers

import (
	"github.com/nospaghetti/crypto-price-api/internal/models"
	"github.com/rs/zerolog"
)

type ChainProvider struct {
	providers []Provider
	logger    *zerolog.Logger
}

func NewChainProvider(providers []Provider, logger *zerolog.Logger) *ChainProvider {
	return &ChainProvider{providers, logger}
}

func (c *ChainProvider) GetHistory() ([]models.Price, error) {
	for _, provider := range c.providers {
		c.logger.Info().Str("provider", provider.GetName()).Msg("Receiving history from provider")
		result, err := provider.GetHistory()

		if err != nil {
			c.logger.Info().Err(err).
				Str("provider", provider.GetName()).
				Msg("Failed to receive history from provider")
			continue
		}

		if len(result) > 0 {
			c.logger.Info().Str("provider", provider.GetName()).Msg("Successfully received history from provider")
			return result, nil
		}
	}

	c.logger.Info().Msg("Failed to receive history from any provider")
	return nil, nil
}

func (c *ChainProvider) GetPrices() ([]models.Price, error) {
	for _, provider := range c.providers {
		result, err := provider.GetPrices()

		if err != nil {
			c.logger.Info().Err(err).
				Str("provider", provider.GetName()).
				Msg("Failed to get prices from provider")
			continue
		}

		if len(result) > 0 {
			return result, nil
		}
	}

	return nil, nil
}

func (c *ChainProvider) GetName() string {
	return "ChainProvider"
}
