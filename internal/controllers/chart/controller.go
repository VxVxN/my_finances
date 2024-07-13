package chart

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Controller struct {
	jwtSecretKey    []byte
	orderCollection *mongo.Collection
}

func Init(client *mongo.Client, jwtSecretKey []byte) *Controller {
	orderCollection := client.Database("myFinances").Collection("orders")
	client.Database("myFinances").CreateCollection(context.TODO(), "orders", options.CreateCollection())
	return &Controller{
		jwtSecretKey:    jwtSecretKey,
		orderCollection: orderCollection,
	}
}
