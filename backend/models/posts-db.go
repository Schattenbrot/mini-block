package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBModel
type DBModel struct {
	DB *mongo.Database
}

func (m *DBModel) InsertPost(post Post) (*primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	currentTime := time.Now().UTC()
	post.CreatedAt = &currentTime
	post.UpdatedAt = &currentTime

	collection := m.DB.Collection("posts")

	oid, err := collection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}

	result := oid.InsertedID.(primitive.ObjectID)

	return &result, nil
}

// GetOnePost returns one post and error, if any
func (m *DBModel) FindOnePost(id primitive.ObjectID) (*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	collection := m.DB.Collection("posts")

	var post Post

	filter := Post{ID: id}
	err := collection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// GetAllPosts returns all posts and error, if any
func (m *DBModel) FindAllPosts() ([]*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := m.DB.Collection("posts")

	findOptions := *options.Find()
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(ctx, bson.D{}, &findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []*Post

	for cursor.Next(ctx) {
		var post Post
		cursor.Decode(&post)
		posts = append(posts, &post)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (m *DBModel) DeleteOnePost(id primitive.ObjectID) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	collection := m.DB.Collection("posts")

	result, err := collection.DeleteOne(ctx, Post{ID: id})
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}
