package services

import "github.com/nospaghetti/crypto-price-api/internal/data/providers"

type HistoryService struct {
	provider providers.HistoryProvider
}

func NewHistoryService(provider providers.HistoryProvider) *HistoryService {
	return &HistoryService{provider}
}

func (s *HistoryService) GetHistory() {
}
