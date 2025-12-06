package services

import (
	"github.com/nospaghetti/crypto-price-api/internal/data/providers"
	gocache "github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
)

type PricesService struct {
	provider providers.Provider
	logger   *zerolog.Logger
	cache    *gocache.Cache
}

func NewPriceService(provider providers.Provider, logger *zerolog.Logger, cache *gocache.Cache) *PricesService {
	return &PricesService{provider, logger, cache}
}

func (s *PricesService) GetPrices(symbol string, currencies []string) (map[string]float64, error) {
	prices, ok := s.cache.Get(symbol)

	if !ok {
		prices, err := s.provider.GetPrices(symbol)

		if err != nil {
			return nil, err
		}

		s.cache.Set(symbol, prices, gocache.DefaultExpiration)
	}

	var output map[string]float64

	for _, currency := range currencies {
		if price, ok := prices.(map[string]float64)[currency]; ok {
			output[currency] = price
		}
	}

	return output, nil
}
