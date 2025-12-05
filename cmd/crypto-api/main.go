// @title Crypto Price API
// @version 1.0
// @description API for fetching crypto currency prices.

// @BasePath  /api/v1

package main

import (
	"context"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nospaghetti/crypto-price-api/internal/app"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "crypto-price-api").
		Logger()

	logger.Info().Msg("Connecting to database...")

	db, err := pgxpool.New(context.Background(), "postgres://user:pass@localhost:5432/mydb")
	if err != nil {
		logger.Fatal().Err(err).Msg("Unable to connect to database")
	}
	defer db.Close()

	logger.Info().Msg("Successfully connected to database")
	logger.Info().Msg("Setting up server...")

	a := app.NewApp(db, &logger)
	mux := http.NewServeMux()
	a.SetupRoutes(mux)
	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		logger.Fatal().Err(err).Msg("failed to start server")
		return
	}

	logger.Info().Msg("Server listening on port 8080")
}
