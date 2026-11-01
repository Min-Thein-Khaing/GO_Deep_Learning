package user

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Repo interface {
	Create(ctx context.Context, user User) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	FindByID(ctx context.Context, id string) (User, error)
	UpdateRefreshToken(ctx context.Context, id string, refreshToken string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, user User) (User, error) {
	record := UserRecord{
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Role:           user.Role,
		RefreshToken:   user.RefreshToken,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
	if err := r.db.WithContext(ctx).Create(&record).Error; err != nil {
		return User{}, fmt.Errorf("insert user failed: %w", err)
	}
	return toUser(record), nil
}

func (r *repo) FindByEmail(ctx context.Context, email string) (User, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	var record UserRecord
	err := r.db.WithContext(ctx).
		Where("LOWER(email) = ?", email).
		First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, fmt.Errorf("user not found: %w", err)
		}
		return User{}, fmt.Errorf("find user failed: %w", err)
	}
	return toUser(record), nil
}

func (r *repo) FindByID(ctx context.Context, id string) (User, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return User{}, fmt.Errorf("invalid user id: %w", err)
	}

	var record UserRecord
	err = r.db.WithContext(ctx).First(&record, intID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, fmt.Errorf("user not found: %w", err)
		}
		return User{}, fmt.Errorf("find user by id failed: %w", err)
	}
	return toUser(record), nil
}

func (r *repo) UpdateRefreshToken(ctx context.Context, id string, refreshToken string) error {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid user id: %w", err)
	}

	result := r.db.WithContext(ctx).Model(&UserRecord{}).
		Where("id = ?", intID).
		Updates(map[string]interface{}{
			"refresh_token": refreshToken,
			"updated_at":    time.Now().UTC(),
		})
	if result.Error != nil {
		return fmt.Errorf("failed to update refresh token: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func toUser(record UserRecord) User {
	return User{
		ID:             strconv.FormatInt(record.ID, 10),
		Email:          record.Email,
		HashedPassword: record.HashedPassword,
		Role:           record.Role,
		RefreshToken:   record.RefreshToken,
		CreatedAt:      record.CreatedAt,
		UpdatedAt:      record.UpdatedAt,
	}
}
