package controllers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Controller struct {
	authCollection          *mongo.Collection
	refreshTokensCollection *mongo.Collection
}

func Init(client *mongo.Client) (*Controller, error) {
	authCollection := client.Database("myFinances").Collection("auth")
	refreshTokensCollection := client.Database("myFinances").Collection("refresh_tokens")

	ctx := context.Background()

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	if _, err := authCollection.Indexes().CreateOne(ctx, indexModel); err != nil {
		return nil, err
	}

	return &Controller{authCollection: authCollection, refreshTokensCollection: refreshTokensCollection}, nil
}
