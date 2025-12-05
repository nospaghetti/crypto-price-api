package providers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nospaghetti/crypto-price-api/internal/models"
	"github.com/rs/zerolog"
)

type CoinGecko struct {
	logger       *zerolog.Logger
	client       *http.Client
	accessHeader string
	accessKey    string
}

func NewCoinGecko(client *http.Client, logger *zerolog.Logger, accessHeader string, accessKey string) *CoinGecko {
	return &CoinGecko{logger, client, accessHeader, accessKey}
}

func (c *CoinGecko) GetPrices() ([]models.Price, error) {
	return []models.Price{}, nil
}

func (c *CoinGecko) GetHistory() ([]models.Price, error) {
	c.logger.Info().Str("provider", c.GetName()).Msg("Sending GetHistory request")
	req, err := c.newRequest("GET", "https://pro-api.coingecko.com/api/v3/coins/markets")
	if err != nil {
		c.logger.Error().Err(err).Str("provider", c.GetName()).Msg("Failed to prepare request")
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		c.logger.Error().Err(err).Str("provider", c.GetName()).Msg("Failed to send request")
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error().Err(err).Str("provider", c.GetName()).Msg("Failed to read response")
		return nil, err
	}

	var raw []map[string]interface{}
	if err := json.Unmarshal(body, &raw); err != nil {
		c.logger.Error().Err(err).Str("provider", c.GetName()).Msg("Failed to unmarshal response")
		return nil, err
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

	c.logger.Info().Str("provider", c.GetName()).Msg("Successfully parsed GetHistory response")

	return prices, nil
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
