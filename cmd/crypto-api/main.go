// @title Crypto Price API
// @version 1.0
// @description API for fetching cryptocurrency prices.

// @BasePath  /api/v1

package main

import (
	"net/http"
	"os"

	"github.com/nospaghetti/crypto-price-api/internal/app"
	"github.com/nospaghetti/crypto-price-api/internal/config"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "crypto-price-api").
		Logger()

	logger.Info().Msg("Loading configuration...")
	cfg, err := config.Load(&logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load configuration")
		return
	}

	logger.Info().Msg("Starting server...")
	a := app.NewApp(cfg, &logger)
	mux := http.NewServeMux()
	a.SetupRoutes(mux)
	err = http.ListenAndServe(":"+os.Getenv("PORT"), mux)

	if err != nil {
		logger.Fatal().Err(err).Msg("failed to start server")
		return
	}

	logger.Info().Msg("Server listening on port 8080")
}
