package main

import (
	"context"
	"fmt"
	"log"
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
	fmt.Println("Connecting to database...")

	db, err := pgxpool.New(context.Background(), "postgres://user:pass@localhost:5432/mydb")
	if err != nil {
		logger.Fatal().Err(err).Msg("Unable to connect to database")
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	logger.Info().Msg("Successfully connected to database")
	logger.Info().Msg("Setting up server...")
	fmt.Println("Successfully connected to database")
	fmt.Println("Setting up server...")

	a := app.NewApp(db, logger)
	mux := http.NewServeMux()
	a.SetupRoutes(mux)

	err = http.ListenAndServe(":80", mux)

	if err != nil {
		logger.Fatal().Err(err).Msg("failed to start server")
		_ = fmt.Errorf("failed to start server: %v", err)

		return
	}

	logger.Info().Msg("Server listening on port 80")
	fmt.Println("Server listening on port 80")
}
