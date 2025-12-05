package services

import (
	"github.com/nospaghetti/crypto-price-api/internal/data/cache"
	"github.com/nospaghetti/crypto-price-api/internal/data/providers"
	"github.com/rs/zerolog"
)

type PricesService struct {
	provider providers.Provider
	logger   *zerolog.Logger
	cache    cache.Cache
}

func NewPriceService(provider providers.Provider, logger *zerolog.Logger, cache cache.Cache) *PricesService {
	return &PricesService{provider, logger, cache}
}

func (s *PricesService) GetPrices(symbol string, currencies []string) (map[string]float64, error) {
	prices, err := s.cache.Get(symbol)

	if err != nil {
		prices, err = s.provider.GetPrices(symbol)

		if err != nil {
			return nil, err
		}

		s.cache.Set(symbol, prices)
	}

	var output map[string]float64

	for _, currency := range currencies {
		if price, ok := prices[currency]; ok {
			output[currency] = price
		}
	}

	return output, nil
}
