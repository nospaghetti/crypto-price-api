package config

type AppConfig struct {
	HTTPPort         string
	EnabledProviders []struct {
		Provider string
		Priority int
	}
	CoinGecko CoinGecko
}
