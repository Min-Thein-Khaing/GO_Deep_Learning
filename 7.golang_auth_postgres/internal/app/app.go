package app

import (
	"context"
	"database/sql"
	"go_auth_postgres/internal/config"
	"go_auth_postgres/internal/db"
)

type App struct {
	Config *config.Config
	DB     *sql.DB
}

func New(ctx context.Context) (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	pgDB, err := db.Connect(cfg)
	if err != nil {
		return nil, err
	}

	return &App{Config: &cfg, DB: pgDB.DB}, nil
}

func (a *App) Close() error {
	if a.DB != nil {
		return a.DB.Close()
	}
	return nil
}

