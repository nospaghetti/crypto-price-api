package app

import "net/http"

func (a *App) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/health", a.V1.Health.Health())
	mux.HandleFunc("/api/v1/prices", a.V1.Prices.GetPrices())
	mux.HandleFunc("/api/v1/history/", a.V1.History.GetHistory())
}
