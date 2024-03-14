package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoConnection struct {
	collection *mongo.Collection
}

func NewMongoConnection(connectionString string) (*mongoConnection, error) {

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
	}, nil
}

func (mc *mongoConnection) Create(entity interface{}) (string, error) {
	_, err := mc.collection.InsertOne(context.Background(), entity)
	if err != nil {
		return "", err
	}
	return "", err
}
func (mc *mongoConnection) Read(slug string, entity interface{}) (interface{}, error) {
	result := mc.collection.FindOne(context.Background(), bson.M{"slug": slug})
	if err := result.Decode(entity); err != nil {
		return nil, err
	}

	return entity, nil
}
func (mc *mongoConnection) Update(entity interface{}) (interface{}, error) {
	return nil, nil
}
func (mc *mongoConnection) Delete(slug string, entity interface{}) error {
	return nil
}
func (mc *mongoConnection) List(entity interface{}) ([]byte, error) {
	return nil, nil
}
func (mc *mongoConnection) ListByTag(tag string, entity interface{}) ([]byte, error) {
	return nil, nil
}
func (mc *mongoConnection) ListByCategory(category string, entity interface{}) ([]byte, error) {
	return nil, nil
}
func (mc *mongoConnection) ListByAuthor(author string, entity interface{}) ([]byte, error) {
	return nil, nil
}
func (mc *mongoConnection) ListByDate(date string, entity interface{}) ([]byte, error) {
	return nil, nil
}
