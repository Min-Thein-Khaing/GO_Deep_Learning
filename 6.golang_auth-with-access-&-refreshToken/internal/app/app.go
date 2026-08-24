package app

import (
	"context"
	"go_auth/internal/config"
	"go_auth/internal/db"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Config *config.Config
	DB     *mongo.Database
	Client *mongo.Client
}

func New(ctx context.Context) (*App, error) {

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	mongoDB, err := db.Connect(cfg)
	if err != nil {
		return nil, err
	}
	return &App{Config: &cfg, DB: mongoDB.DB, Client: mongoDB.Client}, nil

}

func (a *App) Close() error {
	if a.Client != nil {
		return db.Disconnect(&db.MongoDB{Client: a.Client, DB: a.DB})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.Client.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}
