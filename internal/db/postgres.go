package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPostgresDB creates a new database connection
func NewPostgresDB(ctx context.Context, postgresURL string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(ctx, postgresURL)

	if err != nil {
		return nil, err
	}

	if err := dbpool.Ping(ctx); err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL!")
	return dbpool, nil
}
