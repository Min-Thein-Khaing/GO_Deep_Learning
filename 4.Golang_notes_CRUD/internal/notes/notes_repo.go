package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Repo-data access layer

// this is including *mongo.Collection crud method logic
type Repo struct {
	coll *mongo.Collection
}

// constructor function
func NewRepo(db *mongo.Database) *Repo {

	return &Repo{
		coll: db.Collection("notes"),
	}
}

func (r *Repo) Create(ctx context.Context, note Note) (Note, error) {

	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	_, err := r.coll.InsertOne(opCtx, note)

	if err != nil {
		return Note{}, fmt.Errorf("insert note failed: %w", err)
	}

	return note, nil
}

func (r *Repo) List(ctx context.Context) ([]Note, error) {

	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{} //match all docs

	cursor, err := r.coll.Find(opCtx, filter)
	if err != nil {
		return nil, fmt.Errorf("find notes failed: %w", err)
	}
	//cursor must be close after use

	//avoid any kind of memory leak, so we use defer to close the cursor after the function returns
	defer cursor.Close(opCtx)

	var notes []Note
	if err = cursor.All(opCtx, &notes); err != nil {
		return nil, fmt.Errorf("decode notes failed: %w", err)
	}

	return notes, nil
}

func (r *Repo) FindByID(ctx context.Context, id primitive.ObjectID) (Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	var note Note
	err := r.coll.FindOne(opCtx, filter).Decode(&note)
	if err != nil {
		return Note{}, fmt.Errorf("find note failed: %w", err)
	}

	return note, nil
}

func (r *Repo) Delete(ctx context.Context, id primitive.ObjectID) (Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	var note Note
	err := r.coll.FindOneAndDelete(opCtx, filter).Decode(&note)
	if err != nil {
		return Note{}, fmt.Errorf("delete note failed: %w", err)
	}

	return note, nil
}

func (r *Repo) Update(ctx context.Context, id primitive.ObjectID, updatedNote UpdateNoteRequest) (Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{
		"title":     updatedNote.Title,
		"content":   updatedNote.Content,
		"pinned":    updatedNote.Pinned,
		"updatedAt": time.Now().UTC(),
	}}

	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	var note Note
	err := r.coll.FindOneAndUpdate(opCtx, filter, update, &opts).Decode(&note)
	if err != nil {
		return Note{}, fmt.Errorf("update note failed: %w", err)
	}
	return note, nil
}
