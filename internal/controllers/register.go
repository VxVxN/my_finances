package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/VxVxN/my_finances/pkg/httptools"
)

type RegisterRequest struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

type RegisterResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (ctrl *Controller) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := httptools.UnmarshalRequest(r.Body, &req); err != nil {
		httptools.ErrResponse(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)

	if req.Password == "" || req.Username == "" {
		httptools.ErrResponse(w, http.StatusBadRequest, fmt.Errorf("empty username or password"))
		return
	}

	_, err := ctrl.authCollection.InsertOne(context.Background(), req)
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't insert register data: %v", err))
		return
	}

	accessToken, refreshToken, err := ctrl.login(req.Username)
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("failed to login: %v", err))
		return
	}

	httptools.SuccessResponse(w, RegisterResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
