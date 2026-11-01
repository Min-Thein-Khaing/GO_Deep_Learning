package db

import (
	"context"
	"fmt"
	"time"

	"go_auth_mysql/internal/config"
	"go_auth_mysql/internal/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDB struct {
	DB *gorm.DB
}

func Connect(cfg config.Config) (*MySQLDB, error) {
	db, err := gorm.Open(mysql.Open(cfg.MySQLDSN()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get MySQL connection pool: %w", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("failed to ping MySQL: %w", err)
	}

	if err := db.WithContext(ctx).AutoMigrate(&user.UserRecord{}); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("failed to initialize database schema: %w", err)
	}

	fmt.Println("Connected to MySQL")
	return &MySQLDB{DB: db}, nil
}

func Disconnect(mysqlDB *MySQLDB) error {
	if mysqlDB == nil || mysqlDB.DB == nil {
		return nil
	}
	sqlDB, err := mysqlDB.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get MySQL connection pool: %w", err)
	}
	return sqlDB.Close()
}
