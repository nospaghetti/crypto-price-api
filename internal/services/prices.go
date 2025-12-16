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

func (s *PricesService) GetPrices(symbol string, fiats []string) (map[string]float64, error) {
	prices, ok := s.cache.Get(symbol)

	if !ok {
		fetchedPrices, err := s.provider.GetPrices(symbol)

		if err != nil {
			return nil, err
		}

		s.cache.Set(symbol, fetchedPrices, gocache.DefaultExpiration)
		prices = fetchedPrices
	}

	var output = make(map[string]float64, len(fiats))
	for _, currency := range fiats {
		if price, ok := prices.(map[string]float64)[currency]; ok {
			output[currency] = price
		}
	}

	return output, nil
}
