package providers

type Provider interface {
	GetHistory() (map[string]float64, error)
	GetPrices(symbol string) (map[string]float64, error)
	GetName() string
}
