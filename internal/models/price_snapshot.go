package models

type PriceSnapshot struct {
	Symbol    string
	Prices    map[string]float64
	Providers map[string]string
}
