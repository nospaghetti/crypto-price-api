package v1

type GetPricesRequest struct {
	Symbols    []string `json:"symbols"`
	Currencies []string `json:"currencies"`
}

type GetPricesResponse struct {
	Prices    map[string]map[string]float64 `json:"prices"`
	FromCache bool                          `json:"from_cache"`
}
