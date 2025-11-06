package healthcheck

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBChecker struct {
	db *pgxpool.Pool
}

func NewDBChecker(db *pgxpool.Pool) *DBChecker {
	return &DBChecker{db}
}

func (c *DBChecker) Check() error {
	return c.db.Ping(context.Background())
}
