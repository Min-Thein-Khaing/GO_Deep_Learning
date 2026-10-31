package db

import (
	"context"
	"database/sql"
	"fmt"
	"go_auth_postgres/internal/config"
	"time"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	DB *sql.DB
}

func Connect(cfg config.Config) (*PostgresDB, error) {
	connStr := cfg.PostgresDSN()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open Postgres connection: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping Postgres: %w", err)
	}

	fmt.Println("Connected to PostgreSQL")

	if err := initSchema(ctx, db); err != nil {
		return nil, fmt.Errorf("failed to initialize database schema: %w", err)
	}

	return &PostgresDB{DB: db}, nil
}

func initSchema(ctx context.Context, db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id BIGSERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		hashed_password TEXT NOT NULL,
		role VARCHAR(50) NOT NULL DEFAULT 'user',
		refresh_token TEXT,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);`
	_, err := db.ExecContext(ctx, query)
	return err
}

func Disconnect(pgDB *PostgresDB) error {
	if pgDB == nil || pgDB.DB == nil {
		return nil
	}
	return pgDB.DB.Close()
}

