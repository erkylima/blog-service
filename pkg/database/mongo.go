package database

import (
	"context"

	"github.com/erkylima/posts-service/internal/core/ports"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoConnection struct {
	collection *mongo.Collection
}

func NewMongoConnection(connectionString string, repo *ports.Repository) (*mongoConnection, error) {

	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	collection := client.Database("posts_service").Collection("pages")

	return &mongoConnection{
		collection: collection,
		repo:       repo,
	}, nil
}

func (mc *mongoConnection) Push(entity interface{}) error {
	_, err := mc.collection.InsertOne(context.Background(), entity)
	if err != nil {
		return err
	}
	return nil
}
