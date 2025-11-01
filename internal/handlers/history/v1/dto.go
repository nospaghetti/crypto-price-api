package v1

import "time"

type GetHistoryRequest struct {
	Symbols    []string  `json:"symbols"`
	Currencies []string  `json:"currencies"`
	Time       time.Time `json:"time"`
}

type GetHistoryResponse struct {
	Time      time.Time                     `json:"time"`
	Prices    map[string]map[string]float64 `json:"prices"`
	FromCache bool                          `json:"from_cache"`
}
