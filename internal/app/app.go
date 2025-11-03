package app

import (
	v1health "github.com/nospaghetti/crypto-price-api/internal/handlers/health/v1"
	v1history "github.com/nospaghetti/crypto-price-api/internal/handlers/history/v1"
	v1prices "github.com/nospaghetti/crypto-price-api/internal/handlers/prices/v1"
)

type App struct {
	V1 struct {
		Health  *v1health.HandlerV1
		Prices  *v1prices.HandlerV1
		History *v1history.HandlerV1
	}
}

func NewApp() *App {

	return &App{}
}
