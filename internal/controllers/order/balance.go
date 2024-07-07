package order

import (
	"context"
	"fmt"
	"github.com/VxVxN/my_finances/pkg/httptools"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

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
	httptools.SuccessResponse(w, balances)
}
