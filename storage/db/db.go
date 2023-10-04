package db

import (
	"context"
	"fmt"
	"playground/cpp-bootcamp/storage"

	"playground/cpp-bootcamp/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

type strg struct {
	db   *pgxpool.Pool
	user *userRepo
}

func NewStorage(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase,
		),
	)

	if err != nil {
		fmt.Println("ParseConfig:", err.Error())
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		fmt.Println("ConnectConfig:", err.Error())
		return nil, err
	}

	return &strg{
		db: pool,
	}, nil
}

func (d *strg) User() storage.UsersI {
	if d.user == nil {
		d.user = NewUser(d.db)
	}
	return d.user
}
