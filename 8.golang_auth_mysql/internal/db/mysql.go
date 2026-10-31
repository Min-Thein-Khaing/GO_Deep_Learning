package db

import (
	"context"
	"database/sql"
	"fmt"
	"go_auth_mysql/internal/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	DB *sql.DB
}

func Connect(cfg config.Config) (*MySQLDB, error) {
	dsn := cfg.MySQLDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL connection: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %w", err)
	}

	fmt.Println("Connected to MySQL")

	if err := initSchema(ctx, db); err != nil {
		return nil, fmt.Errorf("failed to initialize database schema: %w", err)
	}

	return &MySQLDB{DB: db}, nil
}

func initSchema(ctx context.Context, db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		hashed_password TEXT NOT NULL,
		role VARCHAR(50) NOT NULL DEFAULT 'user',
		refresh_token TEXT,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`
	_, err := db.ExecContext(ctx, query)
	return err
}

func Disconnect(mysqlDB *MySQLDB) error {
	if mysqlDB == nil || mysqlDB.DB == nil {
		return nil
	}
	return mysqlDB.DB.Close()
}

