package cache

import (
	"errors"
	"time"

	"github.com/nospaghetti/crypto-price-api/internal/models"
	gocache "github.com/patrickmn/go-cache"
)

type InMemory struct {
	ttl   time.Duration
	cache *gocache.Cache
}

func NewInMemory(ttl time.Duration, cache *gocache.Cache) *InMemory {
	return &InMemory{ttl, cache}
}

func (i *InMemory) Get(key string) (*models.Price, error) {
	item, success := i.cache.Get(key)

	if !success {
		return nil, errors.New("key not found")
	}

	return item.(*models.Price), nil
}

func (i *InMemory) Set(key string, value models.Price, ttl time.Duration) {
	i.cache.Set(key, value, ttl)
}
