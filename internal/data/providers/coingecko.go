package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/rs/zerolog"
)

type CoinGecko struct {
	logger       *zerolog.Logger
	client       *http.Client
	coinList     map[string]string
	accessHeader string
	accessKey    string
	mutex        *sync.RWMutex
}

func NewCoinGecko(client *http.Client, logger *zerolog.Logger, coinList map[string]string, accessHeader string, accessKey string, mutex *sync.RWMutex) *CoinGecko {
	return &CoinGecko{logger, client, coinList, accessHeader, accessKey, mutex}
}

func (c *CoinGecko) GetPrices(symbol string) (map[string]float64, error) {
	id, ok := c.coinList[symbol]
	log := c.logger.With().
		Str("provider", c.GetName()).
		Str("request", "GetPrices").
		Logger()

	if !ok {
		c.logger.Warn().Str("symbol", symbol).Msg("Symbol not found in configured coin list")
		return nil, errors.New("symbol not found in configured coin list")
	}

	c.logger.Info().Msg("Preparing request")
	req, err := c.newRequest("GET", "https://api.coingecko.com/api/v3/coins/"+id)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to prepare request")
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to send request")
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Error().Int("status_code", resp.StatusCode).Msg("Unexpected status code")
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to read response")
		return nil, err
	}

	var raw struct {
		MarketData struct {
			CurrentPrice map[string]float64 `json:"current_price"`
		} `json:"market_data"`
	}
	if err := json.Unmarshal(body, &raw); err != nil {
		c.logger.Error().Err(err).Msg("Failed to unmarshal response")
		return nil, err
	}

	var prices = make(map[string]float64, len(raw.MarketData.CurrentPrice))
	for s, p := range raw.MarketData.CurrentPrice {
		prices[s] = p
	}

	c.logger.Info().Msg("Successfully parsed response")

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
