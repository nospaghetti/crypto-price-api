package app

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nospaghetti/crypto-price-api/internal/data/cache"
	"github.com/nospaghetti/crypto-price-api/internal/data/providers"
	v1handlers "github.com/nospaghetti/crypto-price-api/internal/handlers/v1"
	"github.com/nospaghetti/crypto-price-api/internal/healthcheck"
	"github.com/nospaghetti/crypto-price-api/internal/services"
	gocache "github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
)

type App struct {
	DB *pgxpool.Pool

	V1 struct {
		Health  *v1handlers.HealthHandler
		Prices  *v1handlers.PricesHandler
		History *v1handlers.HistoryHandler
	}
}

func NewApp(DB *pgxpool.Pool, logger *zerolog.Logger) *App {
	checkers := []healthcheck.Checker{healthcheck.NewDBChecker(DB)}
	client := http.Client{}
	p := []providers.Provider{providers.NewCoinGecko(&client, logger, "x-cg-demo-api-key", "CG-nWzg3BN9TLZdJEs3mNn5eWPA")}
	chainProvider := providers.NewChainProvider(p, logger)
	c := cache.NewInMemory(gocache.New(60*time.Second, 10*time.Minute))

	healthService := services.NewHealthService(checkers)
	pricesService := services.NewPriceService(chainProvider, logger, c)
	historyService := services.NewHistoryService(chainProvider, logger)

	healthHandler := v1handlers.NewHealthHandler(healthService)
	pricesHandler := v1handlers.NewPricesHandler(pricesService)
	historyHandler := v1handlers.NewHistoryHandler(historyService)

	return &App{
		DB: DB,
		V1: struct {
			Health  *v1handlers.HealthHandler
			Prices  *v1handlers.PricesHandler
			History *v1handlers.HistoryHandler
		}{
			Health:  healthHandler,
			Prices:  pricesHandler,
			History: historyHandler,
		},
	}
}
