package controllers

import (
	"fmt"
	"github.com/VxVxN/my_finances/pkg/httptools"
	"net/http"
)

func (ctrl *Controller) RemoveOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" || id == "0" {
		httptools.ErrResponse(w, http.StatusBadRequest, fmt.Errorf("id should not be empty and more than zero: id: %s", id))
		return
	}

	httptools.SuccessResponse(w, nil)
}
