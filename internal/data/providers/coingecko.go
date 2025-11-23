package providers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nospaghetti/crypto-price-api/internal/models"
)

type CoinGecko struct {
	client       *http.Client
	accessHeader string
	accessKey    string
}

func NewCoinGecko(client *http.Client, accessHeader string, accessKey string) *CoinGecko {
	return &CoinGecko{client, accessHeader, accessKey}
}

func (c *CoinGecko) GetPrices() ([]models.Price, error) {
	return []models.Price{}, nil
}

func (c *CoinGecko) GetHistory() ([]models.Price, error) {
	req, err := c.newRequest("GET", "https://pro-api.coingecko.com/api/v3/coins/markets")
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

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

	return []models.Price{}, nil
}

func (c *CoinGecko) GetName() string {
	return "CoinGecko"
}

func (c *CoinGecko) newRequest(method string, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add(c.accessHeader, c.accessKey)
	return req, nil
}
