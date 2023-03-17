package postgres

import (
	"context"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Url         string
	MaxOpenCons int
	MaxIdleCons int
}

func NewClient(ctx context.Context, config Config) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "pgx", config.Url)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxOpenCons)
	db.SetMaxIdleConns(config.MaxIdleCons)

	return db, nil
}
