package app

import (
	"net/http"

	_ "github.com/nospaghetti/crypto-price-api/internal/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/docs/", httpSwagger.WrapHandler)
	mux.HandleFunc("/api/v1/health/", a.V1.Health.Health())
	mux.HandleFunc("/api/v1/prices", a.V1.Prices.GetPrices())
	mux.HandleFunc("/api/v1/history/", a.V1.History.GetHistory())
}
