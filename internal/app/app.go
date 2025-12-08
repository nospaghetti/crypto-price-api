package app

import (
	"net/http"
	"time"

	"github.com/nospaghetti/crypto-price-api/internal/config"
	"github.com/nospaghetti/crypto-price-api/internal/data/providers"
	v1handlers "github.com/nospaghetti/crypto-price-api/internal/handlers/v1"
	"github.com/nospaghetti/crypto-price-api/internal/healthcheck"
	"github.com/nospaghetti/crypto-price-api/internal/services"
	gocache "github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
)

type App struct {
	V1 struct {
		Health *v1handlers.HealthHandler
		Prices *v1handlers.PricesHandler
	}
}

func NewApp(cfg config.AppConfig, logger *zerolog.Logger) *App {
	checkers := []healthcheck.Checker{healthcheck.NewDBChecker()}
	cache := gocache.New(60*time.Second, 10*time.Minute)

	// Crypto providers
	cgClient := http.Client{Timeout: cfg.CoinGecko.HTTPTimeout * time.Second}
	cgProvider := providers.NewCoinGecko(&cgClient, logger, cfg.CoinGecko)
	providerList := []providers.Provider{cgProvider}
	chainProvider := providers.NewChainProvider(providerList, logger)

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
