package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewSqlStorage(config string) (*pgxpool.Pool, error) {
	pgxpool, err := pgxpool.New(context.Background(), config)
	if err != nil {
		return nil, nil
	}

	return pgxpool, nil
}

func InitStorage(pool *pgxpool.Pool) error {
	err := pool.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("could not initialize database: %w", err)
	}

	return nil
}
