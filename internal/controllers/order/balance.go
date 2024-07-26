package order

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/VxVxN/my_finances/pkg/httptools"
	"github.com/VxVxN/my_finances/pkg/tools"
)

var tokens = []string{
	"0x2260fac5e5542a773aa44fbcfedf7c193bc2c599", // WBTC
	"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", // WETH
	"0xdac17f958d2ee523a2206206994597c13d831ec7", // USDT
	"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", // USDC
	"0x6b175474e89094c44da98b954eedeac495271d0f", // DAI
	"0x1d2f0da169ceb9fc7b3144628db156f3f6c60dbe", // XRP
}

type TokenBalance struct {
	Currency   string  `bson:"currency" json:"currency"`
	Balance    float64 `bson:"balance" json:"balance"`
	BalanceUsd float64 `bson:"balance_usd" json:"balance_usd"`
}

type ResponseBalance struct {
	Balance float64        `bson:"balance" json:"balance"`
	Tokens  []TokenBalance `bson:"tokens" json:"tokens"`
}

func (ctrl *Controller) Balance(w http.ResponseWriter, r *http.Request) {
	username, err := httptools.GetValueFromJwtToken(r, ctrl.jwtSecretKey, "username")
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't get username: %v", err))
		return
	}

	cur, err := ctrl.balanceCollection.Find(context.Background(), bson.D{{"username", username}})
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't get balances: %v", err))
		return
	}
	var balances []Balance
	if err = cur.All(context.Background(), &balances); err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't get all balances: %v", err))
		return
	}
	prices, err := tools.FetchPrices(tokens)
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't fetch prices: %v", err))
		return
	}
	var tokenBalance []TokenBalance
	var commonBalance float64
	for _, balance := range balances {
		price, ok := prices[balance.Currency]
		if !ok {
			httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't find currency: %s", balance.Currency))
			return
		}
		tokenBalance = append(tokenBalance, TokenBalance{
			Balance:    balance.Balance,
			Currency:   balance.Currency,
			BalanceUsd: balance.Balance * price})
		commonBalance += balance.Balance * price
	}
	httptools.SuccessResponse(w, ResponseBalance{Balance: commonBalance, Tokens: tokenBalance})
}
