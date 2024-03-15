package database

import (
	"context"

	"github.com/erkylima/posts-service/internal/core/domains"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoConnection[T any] struct {
	Ctx        context.Context
	collection *mongo.Collection
}

func NewMongoConnection[T any](connectionString, dbName, collectionName string) (*mongoConnection[T], error) {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(connectionString)
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	// Check the connection
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)

	return &mongoConnection[T]{
		Ctx:        ctx,
		collection: collection,
	}, nil
}

func (mc *mongoConnection[T]) Create(entity *T) (string, error) {
	result, err := mc.collection.InsertOne(mc.Ctx, entity)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), err
}
func (mc *mongoConnection[T]) Read(slug string, entity *T) error {
	result := mc.collection.FindOne(mc.Ctx, bson.M{"slug": slug})
	if err := result.Decode(entity); err != nil {
		return err
	}

	return nil
}
func (mc *mongoConnection[T]) Update(slug string, entity *T) error {
	var oldEntity T
	result := mc.collection.FindOne(mc.Ctx, bson.M{"slug": slug})
	if err := result.Decode(oldEntity); err != nil {
		return err
	}
	filter := bson.M{"slug": slug}

	update := bson.M{"$set": entity}

	_, err := mc.collection.UpdateOne(mc.Ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
func (mc *mongoConnection[T]) Delete(slug string) error {
	filter := bson.M{"slug": slug}
	_, err := mc.collection.DeleteOne(mc.Ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
func (mc *mongoConnection[T]) List(filters []domains.Filter) ([]T, error) {
	filtersBson := bson.D{}

	// for _, filter := range filters {
	// 	filtersBson = append(filtersBson, bson.E{Key: filter.Key, Value: filter.Value})
	// }
	cursor, err := mc.collection.Find(mc.Ctx, filtersBson)
	if err != nil {
		return nil, err
	}
	var entities []T

	defer cursor.Close(mc.Ctx)
	for cursor.Next(mc.Ctx) {
		var entity T
		if err := cursor.Decode(&entity); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}
