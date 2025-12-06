package providers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	gocache "github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type CoinGecko struct {
	logger       *zerolog.Logger
	client       *http.Client
	cache        *gocache.Cache
	accessHeader string
	accessKey    string
	mutex        *sync.RWMutex
}

func NewCoinGecko(client *http.Client, logger *zerolog.Logger, cache *gocache.Cache, accessHeader string, accessKey string, mutex *sync.RWMutex) *CoinGecko {
	return &CoinGecko{logger, client, cache, accessHeader, accessKey, mutex}
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

	var raw struct {
		MarketData struct {
			CurrentPrice map[string]float64 `json:"current_price"`
		} `json:"market_data"`
	}
	if err := json.Unmarshal(body, &raw); err != nil {
		c.logger.Error().Err(err).Str("provider", c.GetName()).Msg("Failed to unmarshal response")
		return nil, err
	}

	var prices map[string]float64
	for _, item := range raw.MarketData.CurrentPrice {
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
	coins, ok := c.cache.Get("coingecko.coins.list")
	if ok {
		id, ok := coins.(map[string]string)[symbol]
		if ok {
			return id, nil
		}
	}

	if err := c.refreshCoinList(); err != nil {
		return "", err
	}

	var id string
	coins, ok = c.cache.Get("coingecko.coins.list")
	if ok {
		id, ok = coins.(map[string]string)[symbol]
		if !ok {
			return "", fmt.Errorf("failed to get id by symbol")
		}
	}

	return id, nil
}

func (c *CoinGecko) refreshCoinList() error {
	c.logger.With().Str("provider", c.GetName()).Logger()
	log.Info().Msg("Sending request to get coin list")
	// TODO: вынести baseUrl
	req, err := c.newRequest("GET", "https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		log.Error().Err(err).Msg("Failed to prepare request")
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send request")
		return err
	}

	if resp.StatusCode != http.StatusOK {
		log.Error().
			Int("status_code", resp.StatusCode).
			Msg("Failed to get coin list")
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close response body")
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read response")
		return err
	}

	var raw []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
	}
	if err := json.Unmarshal(body, &raw); err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal response")
		return err
	}

	coins := make(map[string]string, len(raw))
	for _, coin := range raw {
		coins[coin.Symbol] = coin.ID
	}

	c.cache.Set("coingecko.coins.list", coins, time.Minute*15)
	log.Info().Msg("Successfully parsed coin list response")

	return nil
}

func (c *CoinGecko) newRequest(method string, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add(c.accessHeader, c.accessKey)
	return req, nil
}
