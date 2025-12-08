package providers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nospaghetti/crypto-price-api/internal/apperr"
	"github.com/nospaghetti/crypto-price-api/internal/config"
	"github.com/rs/zerolog"
)

type CoinGecko struct {
	logger *zerolog.Logger
	client *http.Client
	cfg    config.CoinGecko
}

func NewCoinGecko(client *http.Client, logger *zerolog.Logger, cfg config.CoinGecko) *CoinGecko {
	return &CoinGecko{logger, client, cfg}
}

func (c *CoinGecko) GetPrices(symbol string) (map[string]float64, error) {
	id, ok := c.cfg.CoinIDList[symbol]
	log := c.logger.With().
		Str("provider", c.GetName()).
		Str("request", "GetPrices").
		Logger()

	if !ok {
		c.logger.Warn().Str("symbol", symbol).Msg("Symbol not found in configured coin list")
		return nil, fmt.Errorf("%w: symbol %q not found in configured coin list", apperr.InternalError, symbol)
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close response body")
		}
	}(resp.Body)
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

	req.Header.Add(c.cfg.AuthHeaderName, c.cfg.AuthHeaderValue)
	return req, nil
}
