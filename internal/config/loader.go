package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func Load(logger *zerolog.Logger) (AppConfig, error) {
	var cfg AppConfig
	var err error
	_ = godotenv.Load()

	// Main config
	cfg.HTTPPort, err = getEnv("HTTP_PORT")
	if err != nil {
		return AppConfig{}, err
	}

	val, err := getEnv("ENABLED_PROVIDERS")
	if err != nil {
		return AppConfig{}, err
	}
	if err := json.Unmarshal([]byte(val), &cfg.EnabledProviders); err != nil {
		return AppConfig{}, err
	}
	for i, p := range cfg.EnabledProviders {
		cfg.EnabledProviders[i].Provider = strings.TrimSpace(p.Provider)
		if p.Provider == "" {
			logger.Warn().Msg("empty provider name found in ENABLED_PROVIDERS, skipping")
			continue
		}
	}
	sort.Slice(cfg.EnabledProviders, func(i, j int) bool {
		return cfg.EnabledProviders[i].Priority > cfg.EnabledProviders[j].Priority
	})

	// CoinGecko config
	cfg.CoinGecko.AuthHeaderName, err = getEnv("CG_AUTH_HEADER_NAME")
	if err != nil {
		return AppConfig{}, err
	}
	cfg.CoinGecko.AuthHeaderValue, err = getEnv("CG_AUTH_HEADER_VALUE")
	if err != nil {
		return AppConfig{}, err
	}
	cfg.CoinGecko.BaseURL, err = getEnv("CG_BASE_URL")
	if err != nil {
		return AppConfig{}, err
	}
	val, err = getEnv("CG_HTTP_TIMEOUT")
	if err != nil {
		return AppConfig{}, err
	}
	cfg.CoinGecko.HTTPTimeout, err = time.ParseDuration(val)
	if err != nil {
		return AppConfig{}, err
	}
	val, err = getEnv("CG_COIN_ID_LIST")
	if err != nil {
		return AppConfig{}, err
	}
	if err := json.Unmarshal([]byte(val), &cfg.CoinGecko.CoinIDList); err != nil {
		return AppConfig{}, err
	}

	return cfg, nil
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("missing required %s environment variable", key)
	}

	return value, nil
}
