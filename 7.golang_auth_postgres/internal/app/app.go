package app

import (
	"context"
	"go_auth_postgres/internal/config"
	"go_auth_postgres/internal/db"

	"gorm.io/gorm"
)

type App struct {
	Config *config.Config
	DB     *gorm.DB
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
		sqlDB, err := a.DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
