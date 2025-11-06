package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nospaghetti/crypto-price-api/internal/app"
)

func main() {
	fmt.Println("Connecting to database...")
	db, err := pgxpool.New(context.Background(), "postgres://user:pass@localhost:5432/mydb")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	fmt.Println("Successfully connected to database")
	fmt.Println("Setting up server...")

	a := app.NewApp(db)
	mux := http.NewServeMux()
	a.SetupRoutes(mux)

	err = http.ListenAndServe(":80", mux)

	if err != nil {
		_ = fmt.Errorf("failed to start server: %v", err)

		return
	}

	fmt.Println("Server listening on port 80")
}
