package services

import "github.com/nospaghetti/crypto-price-api/internal/healthcheck"

type HealthService struct {
	checkers []healthcheck.Checker
}

func NewHealthService(checkers []healthcheck.Checker) *HealthService {
	return &HealthService{checkers}
}

func (s *PricesService) GetHealth() {
}
