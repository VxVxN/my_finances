package order

import (
	"context"
	"fmt"
	"github.com/VxVxN/my_finances/pkg/httptools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type Order struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Type     OrderType          `bson:"type" json:"type"`
	Datetime time.Time          `bson:"datetime" json:"datetime"`
	Currency string             `bson:"currency" json:"currency"`
	Amount   float64            `bson:"amount" json:"amount"`
	Price    float64            `bson:"price" json:"price"`
}

func NewOrder(orderType OrderType, dateTime time.Time, currency string, amount float64, price float64) *Order {
	return &Order{
		Id:       primitive.NewObjectID(),
		Type:     orderType,
		Datetime: dateTime,
		Currency: currency,
		Amount:   amount,
		Price:    price,
	}
}

func (ctrl *Controller) Orders(w http.ResponseWriter, r *http.Request) {
	cur, err := ctrl.orderCollection.Find(context.Background(), bson.D{})
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't get orders: %v", err))
		return
	}
	var orders []Order
	if err = cur.All(context.Background(), &orders); err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't get all orders: %v", err))
		return
	}
	httptools.SuccessResponse(w, orders)
}
