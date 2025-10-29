package v1

type GetPricesRequest struct {
	Symbols    []string `json:"symbols"`
	Currencies []string `json:"currencies"`
}

type GetPricesDTO struct {
	Symbol string             `json:"symbol"`
	Prices map[string]float64 `json:"prices"`
}
