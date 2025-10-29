package providers

import (
	"github.com/rs/zerolog"
)

type ChainProvider struct {
	providers []Provider
	logger    *zerolog.Logger
}

func NewChainProvider(providers []Provider, logger *zerolog.Logger) *ChainProvider {
	return &ChainProvider{providers, logger}
}

func (c *ChainProvider) GetPrices(symbol string) (map[string]float64, error) {
	for _, provider := range c.providers {
		c.logger.Info().Str("provider", provider.GetName()).Msg("Getting prices from provider")
		result, err := provider.GetPrices(symbol)

		if err != nil {
			c.logger.Info().Err(err).
				Str("provider", provider.GetName()).
				Msg("Failed to get prices from provider")
			continue
		}

		if len(result) > 0 {
			c.logger.Info().Str("provider", provider.GetName()).Msg("Successfully got prices from provider")
			return result, nil
		}
	}

	return nil, nil
}

func (c *ChainProvider) GetName() string {
	return "ChainProvider"
}
