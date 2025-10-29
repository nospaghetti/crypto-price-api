package app

import "net/http"

func (a *App) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", a.HealthHandler.Health())
	mux.HandleFunc("/prices", a.PriceHandler.CurrentPrice())
	mux.HandleFunc("/history/", a.PriceHandler.HistoryPrice())
}
