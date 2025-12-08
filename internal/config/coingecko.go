package config

import "time"

type CoinGecko struct {
	AuthHeaderName  string
	AuthHeaderValue string
	HTTPTimeout     time.Duration
	BaseURL         string
	CoinIDList      map[string]string
}
