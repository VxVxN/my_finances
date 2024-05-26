package controllers

import (
	"fmt"
	"net/http"
)

func (ctrl *Controller) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hi there!!</h1>")
}
