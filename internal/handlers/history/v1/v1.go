package v1

import "net/http"

type HandlerV1 struct {
}

func (h *HandlerV1) GetHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
