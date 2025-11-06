package providers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nospaghetti/crypto-price-api/internal/models"
)

type CoinGecko struct {
}

func NewCoinGecko() *CoinGecko {
	return &CoinGecko{}
}

func (c *CoinGecko) GetPrices() []models.Price {
	return []models.Price{}
}

func (c *CoinGecko) GetHistory() []models.Price {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://pro-api.coingecko.com/api/v3/coins/markets", nil)
	req.Header.Add("X-CG-DEMO-API-KEY", "CG-nWzg3BN9TLZdJEs3mNn5eWPA")
	resp, err := client.Do(req)

	if err != nil {
		// TODO: handle error
	}

	body, err := io.ReadAll(resp.Body)
	var raw []map[string]interface{}
	if err := json.Unmarshal(body, &raw); err != nil {
		panic(err)
	}

	var prices []models.Price
	for _, item := range raw {
		symbol, _ := item["symbol"].(string)
		price, _ := item["current_price"].(float64)

		prices = append(prices, models.Price{
			Currency:     symbol,
			ExchangeRate: price,
		})
	}

	return []models.Price{}
}
