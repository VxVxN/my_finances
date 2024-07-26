package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/VxVxN/my_finances/pkg/httptools"
	"github.com/golang-jwt/jwt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type refreshTokens struct {
	Username string    `bson:"username"`
	Token    string    `bson:"token"`
	ExpireAt time.Time `bson:"expire_at"`
}

type AccessToken struct {
	Username string    `bson:"username"`
	Token    string    `bson:"token"`
	ExpireAt time.Time `bson:"expire_at"`
}

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

	accessTokenExpireAt := time.Now().Add(time.Hour * time.Duration(ctrl.config.AccessTokenExpiredHours))
	payload := jwt.MapClaims{
		"exp":      accessTokenExpireAt.Unix(),
		"username": req.Username,
	}

	accessToken, err := ctrl.SignedJwtToken(payload, ctrl.jwtSecretKey)
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't sign access token: %v", err))
		return
	}

	refreshTokenExpireAt := time.Now().Add(time.Hour * time.Duration(ctrl.config.RefreshTokenExpiredHours))
	payload = jwt.MapClaims{
		"exp":      refreshTokenExpireAt.Unix(),
		"username": req.Username,
	}
	refreshToken, err := ctrl.SignedJwtToken(payload, ctrl.jwtSecretKey)
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't sign refresh token: %v", err))
		return
	}
	upsert := true
	_, err = ctrl.refreshTokensCollection.ReplaceOne(context.Background(), bson.D{{"username", req.Username}}, refreshTokens{Username: req.Username, Token: refreshToken, ExpireAt: refreshTokenExpireAt}, &options.ReplaceOptions{Upsert: &upsert})
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't insert refresh token: %v", err))
		return
	}

	httptools.SuccessResponse(w, r, LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}

func (ctrl *Controller) SignedJwtToken(payload jwt.MapClaims, jwtSecretKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	accessToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
