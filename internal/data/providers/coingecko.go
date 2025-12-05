package providers

import (
	"encoding/json"
	"io"
	"net/http"

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

func (c *CoinGecko) GetPrices(symbol string) (map[string]float64, error) {
	id, _ := c.getIdBySymbol(symbol)
	c.logger.Info().Str("provider", c.GetName()).Msg("Sending current prices request")
	req, err := c.newRequest("GET", "https://pro-api.coingecko.com/api/v3/coins/"+id)
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

	var prices map[string]float64
	for _, item := range raw {
		symbol, _ := item["symbol"].(string)
		price, _ := item["current_price"].(float64)

		prices[symbol] = price
	}

	c.logger.Info().Str("provider", c.GetName()).Msg("Successfully parsed GetHistory response")

	return prices, nil
}

func (c *CoinGecko) GetHistory() (map[string]float64, error) {
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

	var prices map[string]float64
	for _, item := range raw {
		symbol, _ := item["symbol"].(string)
		price, _ := item["current_price"].(float64)

		prices[symbol] = price
	}

	c.logger.Info().Str("provider", c.GetName()).Msg("Successfully parsed GetHistory response")

	return prices, nil
}

func (c *CoinGecko) GetName() string {
	return "CoinGecko"
}

func (c *CoinGecko) getIdBySymbol(symbol string) (string, error) {
	c.logger.Info().Str("provider", c.GetName()).Msg("Sending id by symbol request")
	req, err := c.newRequest("GET", "https://pro-api.coingecko.com/api/v3/coins/list")
	if err != nil {
		c.logger.Error().Err(err).Str("provider", c.GetName()).Msg("Failed to prepare request")
		return "", err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		c.logger.Error().Err(err).Str("provider", c.GetName()).Msg("Failed to send request")
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error().Err(err).Str("provider", c.GetName()).Msg("Failed to read response")
		return "", err
	}

	var raw []map[string]interface{}
	if err := json.Unmarshal(body, &raw); err != nil {
		c.logger.Error().Err(err).Str("provider", c.GetName()).Msg("Failed to unmarshal response")
		return "", err
	}

	var id string
	for _, item := range raw {
		s, _ := item["symbol"].(string)

		if s == symbol {
			id, _ = item["id"].(string)
		}
	}

	c.logger.Info().Str("provider", c.GetName()).Msg("Successfully parsed id from symbol response")

	return id, nil
}

func (c *CoinGecko) newRequest(method string, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add(c.accessHeader, c.accessKey)
	return req, nil
}
