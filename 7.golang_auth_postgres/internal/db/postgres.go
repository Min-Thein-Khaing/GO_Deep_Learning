package db

import (
	"context"
	"fmt"
	"time"

	"go_auth_postgres/internal/config"
	"go_auth_postgres/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func Connect(cfg config.Config) (*PostgresDB, error) {
	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open Postgres connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get Postgres connection pool: %w", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("failed to ping Postgres: %w", err)
	}

	if err := db.WithContext(ctx).AutoMigrate(&user.UserRecord{}); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("failed to initialize database schema: %w", err)
	}

	fmt.Println("Connected to PostgreSQL")
	return &PostgresDB{DB: db}, nil
}

func Disconnect(pgDB *PostgresDB) error {
	if pgDB == nil || pgDB.DB == nil {
		return nil
	}
	sqlDB, err := pgDB.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get Postgres connection pool: %w", err)
	}
	return sqlDB.Close()
}
