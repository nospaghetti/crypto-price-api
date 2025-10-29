package providers

type Provider interface {
	GetPrices(symbol string) (map[string]float64, error)
	GetName() string
}
