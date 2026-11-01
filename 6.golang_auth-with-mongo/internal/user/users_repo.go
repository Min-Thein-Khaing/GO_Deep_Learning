package user

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo interface {
	Create(ctx context.Context, user User) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	FindByID(ctx context.Context, id string) (User, error)
	UpdateRefreshToken(ctx context.Context, id string, refreshToken string) error
}

type repo struct {
	col *mongo.Collection
}

func NewRepo(db *mongo.Database) Repo {
	return &repo{
		col: db.Collection("users"),
	}
}

func (r *repo) Create(ctx context.Context, user User) (User, error) {

	res, err := r.col.InsertOne(ctx, user)

	if err != nil {
		return User{}, fmt.Errorf("insert user failed: %w", err)
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return User{}, fmt.Errorf("failed to convert inserted ID to ObjectID")
	}
	user.ID = id.Hex()

	return user, nil
}

func (r *repo) FindByEmail(ctx context.Context, email string) (User, error) {
	email = strings.ToLower(strings.TrimSpace(email)) // Normalize email to lowercase for case-insensitive search
	var user User
	filter := bson.M{"email": email}
	err := r.col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return User{}, fmt.Errorf("find user failed: %w", err)
	}
	return user, nil
}

func (r *repo) FindByID(ctx context.Context, id string) (User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return User{}, fmt.Errorf("invalid user id: %w", err)
	}
	var user User
	filter := bson.M{"_id": objID}
	err = r.col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return User{}, fmt.Errorf("find user by id failed: %w", err)
	}
	return user, nil
}

func (r *repo) UpdateRefreshToken(ctx context.Context, id string, refreshToken string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid user id: %w", err)
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"refresh_token": refreshToken,
			"updatedAt":     time.Now().UTC(),
		},
	}
	_, err = r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update refresh token: %w", err)
	}
	return nil
}


