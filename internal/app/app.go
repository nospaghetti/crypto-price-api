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
		Health *v1handlers.HealthHandler
		Prices *v1handlers.PricesHandler
	}
}

func NewApp(logger *zerolog.Logger) *App {
	checkers := []healthcheck.Checker{healthcheck.NewDBChecker()}
	client := http.Client{}
	cache := gocache.New(60*time.Second, 10*time.Minute)
	mutex := &sync.RWMutex{}

	cgcoinList := map[string]string{
		"btc": "bitcoin",
		"eth": "ethereum",
	}
	cg := providers.NewCoinGecko(&client, logger, cgcoinList, "x-cg-demo-api-key", "CG-nWzg3BN9TLZdJEs3mNn5eWPA", mutex)
	p := []providers.Provider{cg}
	chainProvider := providers.NewChainProvider(p, logger)

	healthService := services.NewHealthService(checkers)
	pricesService := services.NewPriceService(chainProvider, logger, cache)

	healthHandler := v1handlers.NewHealthHandler(healthService)
	pricesHandler := v1handlers.NewPricesHandler(pricesService)

	return &App{
		V1: struct {
			Health *v1handlers.HealthHandler
			Prices *v1handlers.PricesHandler
		}{
			Health: healthHandler,
			Prices: pricesHandler,
		},
	}
}
