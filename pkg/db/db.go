package db

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
	configs "github.com/samurenkoroma/agro-platform/internal/shared/config"
)

func NewDB(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(context.Background()); err != nil {
		return nil, err
	}
	return db, nil
}

func NewPool(config configs.DbConfig) (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), config.Dsn)
}
