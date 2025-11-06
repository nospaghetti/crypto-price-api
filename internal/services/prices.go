package services

import "github.com/nospaghetti/crypto-price-api/internal/data/providers"

type PricesService struct {
	provider providers.PriceProvider
}

func NewPriceService(provider providers.PriceProvider) *PricesService {
	return &PricesService{provider}
}

func (s *PricesService) GetPrices() {
}
