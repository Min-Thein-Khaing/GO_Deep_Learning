package user

import (
	"context"
	"errors"
	"fmt"
	"go_auth/internal/auth"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo Repo

	jwtSecret string
}

func NewService(repo Repo, jwtSecret string) *Service {
	return &Service{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

type RegisterInput struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResult struct {
	AccessToken  string     `json:"access_token"`
	RefreshToken string     `json:"refresh_token"`
	User         PublicUser `json:"user"`
}

func (s *Service) Register(ctx context.Context, input RegisterInput) (AuthResult, error) {

	email := strings.ToLower(strings.TrimSpace(input.Email)) // Normalize email to lowercase for case-insensitive search
	pass := strings.TrimSpace(input.Password)

	if email == "" || pass == "" {
		return AuthResult{}, errors.New("invalid credentials")
	}

	if len(pass) < 6 {
		return AuthResult{}, errors.New("password must be at least 6 characters")
	}

	_, err := s.repo.FindByEmail(ctx, email)
	if err == nil {
		return AuthResult{}, errors.New("user already exists")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return AuthResult{}, fmt.Errorf("hash password failed: %w", err)
	}
	now := time.Now().UTC()
	user := User{
		Email:          email,
		HashedPassword: string(hashedPass),
		Role:           "user",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	user, err = s.repo.Create(ctx, user)
	if err != nil {
		return AuthResult{}, fmt.Errorf("create user failed: %w", err)
	}

	tokens, err := auth.GenerateTokenPair(s.jwtSecret, user.ID, user.Role)
	if err != nil {
		return AuthResult{}, fmt.Errorf("create tokens failed: %w", err)
	}

	return AuthResult{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User:         user.ToPublicUser(),
	}, nil
}

func (s *Service) Login(ctx context.Context, input LoginInput) (AuthResult, error) {

	email := strings.ToLower(strings.TrimSpace(input.Email))
	pass := strings.TrimSpace(input.Password)

	if email == "" || pass == "" {
		return AuthResult{}, errors.New("invalid credentials")
	}

	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return AuthResult{}, errors.New("invalid credentials")
	}

	compared := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(pass))

	if compared != nil {
		return AuthResult{}, errors.New("invalid credentials")
	}

	tokens, err := auth.GenerateTokenPair(s.jwtSecret, user.ID, user.Role)
	if err != nil {
		return AuthResult{}, fmt.Errorf("create tokens failed: %w", err)
	}

	return AuthResult{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User:         user.ToPublicUser(),
	}, nil

}

func (s *Service) RefreshToken(ctx context.Context, refreshTokenStr string) (AuthResult, error) {
	refreshTokenStr = strings.TrimSpace(refreshTokenStr)
	if refreshTokenStr == "" {
		return AuthResult{}, errors.New("refresh token is required")
	}

	claims, err := auth.ParseAndValidateToken(s.jwtSecret, refreshTokenStr)
	if err != nil {
		return AuthResult{}, fmt.Errorf("invalid or expired refresh token: %w", err)
	}

	if claims.TokenType != "" && claims.TokenType != "refresh" {
		return AuthResult{}, errors.New("invalid token type, refresh token expected")
	}

	user, err := s.repo.FindByID(ctx, claims.Subject)
	if err != nil {
		return AuthResult{}, errors.New("user not found")
	}

	tokens, err := auth.GenerateTokenPair(s.jwtSecret, user.ID, user.Role)
	if err != nil {
		return AuthResult{}, fmt.Errorf("generate tokens failed: %w", err)
	}

	return AuthResult{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User:         user.ToPublicUser(),
	}, nil
}

