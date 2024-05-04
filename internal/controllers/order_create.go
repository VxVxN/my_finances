package controllers

import (
	"github.com/VxVxN/my_finances/pkg/httptools"
	"net/http"
)

type CreateOrderRequest struct {
	Type          string `json:"type"`
	Datetime      string `json:"datetime"`
	BaseCurrency  string `json:"base_currency"`
	Value         int    `json:"value"`
	QuoteCurrency string `json:"quote_currency"`
	Price         int    `json:"price"`
}

func (ctrl *Controller) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest

	if err := httptools.UnmarshalRequest(r.Body, &req); err != nil {
		httptools.ErrResponse(w, http.StatusBadRequest, err)
		return
	}
	httptools.SuccessResponse(w, nil)
}
