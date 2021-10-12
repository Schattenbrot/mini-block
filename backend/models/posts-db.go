package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DBModel
type DBModel struct {
	DB *mongo.Database
}

func (m *DBModel) PostPost(post Post) (*primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println(post)
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
func (m *DBModel) GetOnePost(id primitive.ObjectID) (*Post, error) {
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
func (m *DBModel) GetAllPosts() ([]*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := m.DB.Collection("posts")

	cursor, err := collection.Find(ctx, bson.M{})
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
