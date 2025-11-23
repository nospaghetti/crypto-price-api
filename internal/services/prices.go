package services

import (
	"github.com/nospaghetti/crypto-price-api/internal/data/providers"
	"github.com/rs/zerolog"
)

type PricesService struct {
	provider providers.Provider
	logger   *zerolog.Logger
}

func NewPriceService(provider providers.Provider, logger *zerolog.Logger) *PricesService {
	return &PricesService{provider, logger}
}

func (s *PricesService) GetPrices() {
}
