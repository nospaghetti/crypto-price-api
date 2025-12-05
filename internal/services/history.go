package services

import (
	"github.com/nospaghetti/crypto-price-api/internal/data/providers"
	"github.com/rs/zerolog"
)

type HistoryService struct {
	provider providers.Provider
	logger   *zerolog.Logger
}

func NewHistoryService(provider providers.Provider, logger *zerolog.Logger) *HistoryService {
	return &HistoryService{provider, logger}
}

func (s *HistoryService) GetHistory() {
}
