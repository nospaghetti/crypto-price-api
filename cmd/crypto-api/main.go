package main

import (
	"fmt"
	"net/http"

	"github.com/nospaghetti/crypto-price-api/internal/app"
)

func main() {
	a := &app.App{}

	mux := http.NewServeMux()
	a.SetupRoutes(mux)

	err := http.ListenAndServe(":80", mux)

	if err != nil {
		_ = fmt.Errorf("failed to start server: %v", err)

		return
	}

	fmt.Println("Server listening on port 80")
}
