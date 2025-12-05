package cache

type Cache interface {
	Get(key string) (map[string]float64, error)
	Set(symbol string, prices map[string]float64)
}
