package controllers

import (
	"context"
	"fmt"
	"github.com/VxVxN/my_finances/pkg/httptools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (ctrl *Controller) RemoveOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		httptools.ErrResponse(w, http.StatusBadRequest, fmt.Errorf("id should not be empty: id: %s", id))
		return
	}
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		httptools.ErrResponse(w, http.StatusBadRequest, fmt.Errorf("invalid id: %s", id))
		return
	}
	result, err := ctrl.orderCollection.DeleteOne(context.Background(), bson.M{"_id": idPrimitive})
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("could not remove order: %v", err))
		return
	}
	if result.DeletedCount == 0 {
		httptools.ErrResponse(w, http.StatusNotFound, fmt.Errorf("order not found"))
		return
	}
	httptools.SuccessResponse(w, nil)
}
