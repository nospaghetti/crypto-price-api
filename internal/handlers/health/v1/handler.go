package v1

import (
	"encoding/json"
	"net/http"
)

type HandlerV1 struct {
}

func (h *HandlerV1) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"message": "OK",
		}

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	}
}
