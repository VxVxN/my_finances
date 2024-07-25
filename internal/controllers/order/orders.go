package order

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/VxVxN/my_finances/internal/entities/order"
	"github.com/VxVxN/my_finances/pkg/httptools"
)

func (ctrl *Controller) Orders(w http.ResponseWriter, r *http.Request) {
	cur, err := ctrl.orderCollection.Find(context.Background(), bson.D{})
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't get orders: %v", err))
		return
	}
	var orders []order.Order
	if err = cur.All(context.Background(), &orders); err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't get all orders: %v", err))
		return
	}
	httptools.SuccessResponse(w, r, orders)
}
