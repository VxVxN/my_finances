package chart

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	entityorder "github.com/VxVxN/my_finances/internal/entities/order"
	"github.com/VxVxN/my_finances/pkg/httptools"
)

type TokenBalance struct {
	Date       string  `json:"date"`
	BalanceUsd float64 `json:"balanceUsd"`
}

type HistoricalBalanceResponse struct {
	Tokens map[string][]TokenBalance `json:"tokens"`
}

func (ctrl *Controller) HistoricalBalance(w http.ResponseWriter, r *http.Request) {
	username, err := httptools.GetValueFromJwtToken(r, ctrl.jwtSecretKey, "username")
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't get username: %v", err))
		return
	}
	opts := options.Find().SetSort(bson.D{{"datetime", 1}})
	cur, err := ctrl.orderCollection.Find(context.Background(), bson.D{{"username", username}}, opts)
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't get orders: %v", err))
		return
	}
	var orders []entityorder.Order
	if err = cur.All(context.Background(), &orders); err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't get all orders: %v", err))
		return
	}
	resp := HistoricalBalanceResponse{
		Tokens: make(map[string][]TokenBalance),
	}
	balanceByToken := make(map[string]float64)
	for _, order := range orders {
		if _, ok := resp.Tokens[order.Currency]; !ok {
			resp.Tokens[order.Currency] = append(resp.Tokens[order.Currency], TokenBalance{
				Date:       order.Datetime.Format("2006-01-02"),
				BalanceUsd: order.Value(),
			})
			balanceByToken[order.Currency] = order.Value()
			continue
		}
		var found bool
		for i, tokenBalance := range resp.Tokens[order.Currency] {
			if tokenBalance.Date == order.Datetime.Format("2006-01-02") {
				balanceByToken[order.Currency] = order.Value() + balanceByToken[order.Currency]
				resp.Tokens[order.Currency][i].BalanceUsd = balanceByToken[order.Currency]
				found = true
				break
			}
		}
		if !found {
			balanceByToken[order.Currency] = order.Value() + balanceByToken[order.Currency]
			resp.Tokens[order.Currency] = append(resp.Tokens[order.Currency], TokenBalance{
				Date:       order.Datetime.Format("2006-01-02"),
				BalanceUsd: balanceByToken[order.Currency],
			})
		}
	}

	httptools.SuccessResponse(w, r, resp)
}
