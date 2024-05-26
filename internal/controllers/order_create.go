package controllers

import (
	"context"
	"fmt"
	"github.com/VxVxN/my_finances/pkg/httptools"
	"net/http"
	"time"
)

type CreateOrderRequest struct {
	Type          string    `json:"type"`
	Datetime      time.Time `json:"datetime"`
	BaseCurrency  string    `json:"base_currency"`
	Value         float64   `json:"value"`
	QuoteCurrency string    `json:"quote_currency"`
	Price         float64   `json:"price"`
}

func (ctrl *Controller) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest

	if err := httptools.UnmarshalRequest(r.Body, &req); err != nil {
		httptools.ErrResponse(w, http.StatusBadRequest, err)
		return
	}
	newOrder := NewOrder(req.Type, req.Datetime, req.BaseCurrency, req.Value, req.QuoteCurrency, req.Price)
	_, err := ctrl.orderCollection.InsertOne(context.Background(), newOrder)
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't insert order: %v", err))
		return
	}
	httptools.SuccessResponse(w, nil)
}
