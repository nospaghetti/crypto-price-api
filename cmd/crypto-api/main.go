// @title Crypto Price API
// @version 1.0
// @description API for fetching cryptocurrency prices.

// @BasePath  /api/v1

package main

import (
	"net/http"
	"os"

	"github.com/nospaghetti/crypto-price-api/internal/app"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "crypto-price-api").
		Logger()
	logger.Info().Msg("Setting up server...")

	a := app.NewApp(&logger)
	mux := http.NewServeMux()
	a.SetupRoutes(mux)
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		logger.Fatal().Err(err).Msg("failed to start server")
		return
	}

	logger.Info().Msg("Server listening on port 8080")
}
