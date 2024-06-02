package controllers

import (
	"context"
	"fmt"
	"github.com/VxVxN/my_finances/pkg/httptools"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
	"time"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

var JwtSecretKey = []byte(uuid.NewString())

func (ctrl *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := httptools.UnmarshalRequest(r.Body, &req); err != nil {
		httptools.ErrResponse(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal request body: %v", err))
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)

	if req.Password == "" || req.Username == "" {
		httptools.ErrResponse(w, http.StatusBadRequest, fmt.Errorf("empty username or password"))
		return
	}

	result := ctrl.authCollection.FindOne(context.Background(), bson.D{})
	if result.Err() != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("user not found: %v", result.Err()))
		return
	}

	payload := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(), // todo move to config
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	accessToken, err := token.SignedString(JwtSecretKey)
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't sign token: %v", err))
	}

	httptools.SuccessResponse(w, LoginResponse{AccessToken: accessToken})
}
