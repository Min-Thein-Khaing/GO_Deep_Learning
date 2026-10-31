package user

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Repo interface {
	Create(ctx context.Context, user User) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	FindByID(ctx context.Context, id string) (User, error)
	UpdateRefreshToken(ctx context.Context, id string, refreshToken string) error
}

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, user User) (User, error) {
	query := `
	INSERT INTO users (email, hashed_password, role, refresh_token, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?);
	`
	res, err := r.db.ExecContext(ctx, query,
		user.Email,
		user.HashedPassword,
		user.Role,
		user.RefreshToken,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return User{}, fmt.Errorf("insert user failed: %w", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return User{}, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	user.ID = strconv.FormatInt(lastID, 10)
	return user, nil
}

func (r *repo) FindByEmail(ctx context.Context, email string) (User, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	query := `
	SELECT id, email, hashed_password, role, COALESCE(refresh_token, ''), created_at, updated_at
	FROM users
	WHERE LOWER(email) = LOWER(?);
	`
	var id int64
	var user User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&id,
		&user.Email,
		&user.HashedPassword,
		&user.Role,
		&user.RefreshToken,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("user not found: %w", err)
		}
		return User{}, fmt.Errorf("find user failed: %w", err)
	}

	user.ID = strconv.FormatInt(id, 10)
	return user, nil
}

func (r *repo) FindByID(ctx context.Context, id string) (User, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return User{}, fmt.Errorf("invalid user id: %w", err)
	}

	query := `
	SELECT id, email, hashed_password, role, COALESCE(refresh_token, ''), created_at, updated_at
	FROM users
	WHERE id = ?;
	`
	var user User
	err = r.db.QueryRowContext(ctx, query, intID).Scan(
		&intID,
		&user.Email,
		&user.HashedPassword,
		&user.Role,
		&user.RefreshToken,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("user not found: %w", err)
		}
		return User{}, fmt.Errorf("find user by id failed: %w", err)
	}

	user.ID = strconv.FormatInt(intID, 10)
	return user, nil
}

func (r *repo) UpdateRefreshToken(ctx context.Context, id string, refreshToken string) error {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid user id: %w", err)
	}

	query := `
	UPDATE users
	SET refresh_token = ?, updated_at = ?
	WHERE id = ?;
	`
	_, err = r.db.ExecContext(ctx, query, refreshToken, time.Now().UTC(), intID)
	if err != nil {
		return fmt.Errorf("failed to update refresh token: %w", err)
	}

	return nil
}

