package app

import (
	"net/http"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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

func NewApp(logger *zerolog.Logger) *App {
	checkers := []healthcheck.Checker{healthcheck.NewDBChecker()}
	client := http.Client{}
	cache := gocache.New(60*time.Second, 10*time.Minute)
	mutex := &sync.RWMutex{}
	p := []providers.Provider{providers.NewCoinGecko(&client, logger, cache, "x-cg-demo-api-key", "CG-nWzg3BN9TLZdJEs3mNn5eWPA", mutex)}
	chainProvider := providers.NewChainProvider(p, logger)

	healthService := services.NewHealthService(checkers)
	pricesService := services.NewPriceService(chainProvider, logger, cache)
	historyService := services.NewHistoryService(chainProvider, logger)

	healthHandler := v1handlers.NewHealthHandler(healthService)
	pricesHandler := v1handlers.NewPricesHandler(pricesService)
	historyHandler := v1handlers.NewHistoryHandler(historyService)

	return &App{
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
