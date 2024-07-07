package order

import "go.mongodb.org/mongo-driver/mongo"

type Controller struct {
	jwtSecretKey      []byte
	orderCollection   *mongo.Collection
	balanceCollection *mongo.Collection
}

func Init(client *mongo.Client, jwtSecretKey []byte) *Controller {
	orderCollection := client.Database("myFinances").Collection("orders")
	balanceCollection := client.Database("myFinances").Collection("balance")
	return &Controller{
		jwtSecretKey:      jwtSecretKey,
		orderCollection:   orderCollection,
		balanceCollection: balanceCollection,
	}
}
