package services

import (
	"github.com/nospaghetti/crypto-price-api/internal/data/cache"
	"github.com/nospaghetti/crypto-price-api/internal/data/providers"
	"github.com/rs/zerolog"
)

type PricesService struct {
	provider providers.Provider
	logger   *zerolog.Logger
	cache    *cache.Cache
}

func NewPriceService(provider providers.Provider, logger *zerolog.Logger, cache *cache.Cache) *PricesService {
	return &PricesService{provider, logger, cache}
}

func (s *PricesService) GetPrices() {
}
