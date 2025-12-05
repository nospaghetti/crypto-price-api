package cache

import (
	"time"

	"github.com/nospaghetti/crypto-price-api/internal/models"
)

type Cache interface {
	Get(symbol string) (models.Price, error)
	Set(symbol string, price models.Price, ttl time.Duration)
}
