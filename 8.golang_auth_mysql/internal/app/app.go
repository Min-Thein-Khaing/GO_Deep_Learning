package app

import (
	"context"
	"database/sql"
	"go_auth_mysql/internal/config"
	"go_auth_mysql/internal/db"
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

	mysqlDB, err := db.Connect(cfg)
	if err != nil {
		return nil, err
	}

	return &App{Config: &cfg, DB: mysqlDB.DB}, nil
}

func (a *App) Close() error {
	if a.DB != nil {
		return a.DB.Close()
	}
	return nil
}

