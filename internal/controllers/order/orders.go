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
	Id            primitive.ObjectID `bson:"_id" json:"id"`
	Type          string             `bson:"type" json:"type"`
	Datetime      time.Time          `bson:"datetime" json:"datetime"`
	BaseCurrency  string             `bson:"base_currency" json:"base_currency"`
	Value         float64            `bson:"value" json:"value"`
	QuoteCurrency string             `bson:"quote_currency" json:"quote_currency"`
	Price         float64            `bson:"price" json:"price"`
}

func NewOrder(orderType string, dateTime time.Time, baseCurrency string, value float64, quoteCurrency string, price float64) *Order {
	return &Order{
		Id:            primitive.NewObjectID(),
		Type:          orderType,
		Datetime:      dateTime,
		BaseCurrency:  baseCurrency,
		Value:         value,
		QuoteCurrency: quoteCurrency,
		Price:         price,
	}
}

// type ResponseBody struct {
// 	result: []Order `json:"result"`
// }

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
