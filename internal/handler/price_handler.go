package handler

import "net/http"

type PriceHandler struct {
}

func (h *PriceHandler) CurrentPrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *PriceHandler) HistoryPrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
