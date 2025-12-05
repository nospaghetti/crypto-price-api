package cache

import (
	"errors"

	gocache "github.com/patrickmn/go-cache"
)

type InMemory struct {
	cache *gocache.Cache
}

func NewInMemory(cache *gocache.Cache) *InMemory {
	return &InMemory{cache}
}

func (i *InMemory) Get(key string) (map[string]float64, error) {
	item, success := i.cache.Get(key)

	if !success {
		return nil, errors.New("key not found")
	}

	return item.(map[string]float64), nil
}

func (i *InMemory) Set(key string, prices map[string]float64) {
	i.cache.Set(key, prices, 0)
}
