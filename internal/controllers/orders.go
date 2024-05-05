package controllers

import (
	"github.com/VxVxN/my_finances/pkg/httptools"
	"net/http"
	"time"
)

type Order struct {
	Id            int       `json:"id"`
	Type          string    `json:"type"`
	Datetime      time.Time `json:"datetime"`
	BaseCurrency  string    `json:"base_currency"`
	Value         float64   `json:"value"`
	QuoteCurrency string    `json:"quote_currency"`
	Price         float64   `json:"price"`
}

// type ResponseBody struct {
// 	result: []Order `json:"result"`
// }

func (ctrl *Controller) Orders(w http.ResponseWriter, r *http.Request) {
	httptools.SuccessResponse(w, []Order{
		{
			Id:            1,
			Type:          "buy",
			Datetime:      time.Now(),
			BaseCurrency:  "WETH",
			Value:         1,
			QuoteCurrency: "USD",
			Price:         3200,
		},
		{
			Id:            2,
			Type:          "sell",
			Datetime:      time.Now(),
			BaseCurrency:  "WETH",
			Value:         1,
			QuoteCurrency: "USD",
			Price:         3400,
		},
	})
}
