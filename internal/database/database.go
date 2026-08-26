package database

import (
	"context"

	"github.com/herojk64/portfolio-backend/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg *config.Config) (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), cfg.Database.DSN())
}
