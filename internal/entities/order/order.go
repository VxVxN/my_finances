package order

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Type string

const (
	Buy  Type = "buy"
	Sell      = "sell"
)

type Order struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Username string             `bson:"username" json:"username"`
	Type     Type               `bson:"type" json:"type"`
	Datetime time.Time          `bson:"datetime" json:"datetime"`
	Currency string             `bson:"currency" json:"currency"`
	Amount   float64            `bson:"amount" json:"amount"`
	Price    float64            `bson:"price" json:"price"`
}

func NewOrder(username string, orderType Type, dateTime time.Time, currency string, amount float64, price float64) *Order {
	return &Order{
		Id:       primitive.NewObjectID(),
		Username: username,
		Type:     orderType,
		Datetime: dateTime,
		Currency: currency,
		Amount:   amount,
		Price:    price,
	}
}

func (o *Order) Value() float64 {
	if o.Type == Buy {
		return o.Amount * o.Price
	}
	return o.Amount * o.Price * -1
}
