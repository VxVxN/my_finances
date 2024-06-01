package order

import "go.mongodb.org/mongo-driver/mongo"

type Controller struct {
	orderCollection *mongo.Collection
}

func Init(client *mongo.Client) *Controller {
	orderCollection := client.Database("myFinances").Collection("orders")
	return &Controller{orderCollection: orderCollection}
}
