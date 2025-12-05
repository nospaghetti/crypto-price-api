package main

import (
	"fmt"
	"net/http"

	// ...
	_ "github.com/nospaghetti/crypto-price-api/internal/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/docs/", httpSwagger.WrapHandler)
	fmt.Println("Starting server...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on port 8080")
	// print listening
}
