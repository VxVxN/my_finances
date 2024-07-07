package order

import (
	"context"
	"fmt"
	"github.com/VxVxN/my_finances/pkg/httptools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

type CreateOrderRequest struct {
	Type     OrderType `json:"type"`
	Datetime time.Time `json:"datetime"`
	Currency string    `json:"currency"`
	Amount   float64   `json:"amount"`
	Price    float64   `json:"price"`
}

type OrderType string

const (
	Buy  OrderType = "buy"
	Sell           = "sell"
)

type Balance struct {
	Username string  `bson:"username" json:"username"`
	Currency string  `bson:"currency" json:"currency"`
	Balance  float64 `bson:"balance" json:"balance"`
}

func (ctrl *Controller) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest

	if err := httptools.UnmarshalRequest(r.Body, &req); err != nil {
		httptools.ErrResponse(w, http.StatusBadRequest, err)
		return
	}
	newOrder := NewOrder(req.Type, req.Datetime, req.Currency, req.Amount, req.Price)
	_, err := ctrl.orderCollection.InsertOne(context.Background(), newOrder)
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't insert order: %v", err))
		return
	}

	if err = ctrl.UpdateBalance(r, &req); err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't update balance: %v", err))
		return
	}

	httptools.SuccessResponse(w, nil)
}

func (ctrl *Controller) UpdateBalance(r *http.Request, req *CreateOrderRequest) error {
	username, err := httptools.GetValueFromJwtToken(r, ctrl.jwtSecretKey, "username")
	if err != nil {
		return fmt.Errorf("can't get username: %v", err)
	}

	ctx := context.Background()
	filter := bson.M{"username": username, "currency": req.Currency}
	amount := req.Amount
	if req.Type == Sell {
		amount = -amount
	}
	update := bson.M{
		"$inc": bson.M{"balance": amount},
	}
	opts := options.Update().SetUpsert(true)

	if _, err = ctrl.balanceCollection.UpdateOne(ctx, filter, update, opts); err != nil {
		return err
	}

	return nil
}
