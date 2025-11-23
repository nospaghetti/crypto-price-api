package app

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nospaghetti/crypto-price-api/internal/data/providers"
	v1health "github.com/nospaghetti/crypto-price-api/internal/handlers/health/v1"
	v1history "github.com/nospaghetti/crypto-price-api/internal/handlers/history/v1"
	v1prices "github.com/nospaghetti/crypto-price-api/internal/handlers/prices/v1"
	"github.com/nospaghetti/crypto-price-api/internal/healthcheck"
	"github.com/nospaghetti/crypto-price-api/internal/services"
	"github.com/rs/zerolog"
)

type App struct {
	DB *pgxpool.Pool

	V1 struct {
		Health  *v1health.HealthHandler
		Prices  *v1prices.PricesHandler
		History *v1history.HistoryHandler
	}
}

func NewApp(DB *pgxpool.Pool, logger *zerolog.Logger) *App {
	checkers := []healthcheck.Checker{healthcheck.NewDBChecker(DB)}
	client := http.Client{}
	provs := []providers.Provider{providers.NewCoinGecko(&client, "", "")}
	chainProvider := providers.NewChainProvider(provs, logger)

	healthService := services.NewHealthService(checkers)
	pricesService := services.NewPriceService(provider)
	historyService := services.NewHistoryService(provider)

	healthHandler := v1health.NewHealthHandler(healthService)
	pricesHandler := v1prices.NewPricesHandler(pricesService)
	historyHandler := v1history.NewHistoryHandler(historyService)

	return &App{
		DB: DB,
		V1: struct {
			Health  *v1health.HealthHandler
			Prices  *v1prices.PricesHandler
			History *v1history.HistoryHandler
		}{
			Health:  healthHandler,
			Prices:  pricesHandler,
			History: historyHandler,
		},
	}
}
