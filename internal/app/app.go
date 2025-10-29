package app

import "github.com/nospaghetti/crypto-price-api/internal/handler"

type App struct {
	HealthHandler *handler.HealthHandler
	PriceHandler  *handler.PriceHandler
}
