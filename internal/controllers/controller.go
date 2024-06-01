package controllers

import "go.mongodb.org/mongo-driver/mongo"

type Controller struct {
	authCollection *mongo.Collection
}

func Init(client *mongo.Client) *Controller {
	authCollection := client.Database("myFinances").Collection("auth")
	return &Controller{authCollection: authCollection}
}
