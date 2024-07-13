package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/VxVxN/my_finances/pkg/httptools"
	"github.com/golang-jwt/jwt"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (ctrl *Controller) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshTokenRequest
	if err := httptools.UnmarshalRequest(r.Body, &req); err != nil {
		httptools.ErrResponse(w, http.StatusBadRequest, fmt.Errorf("can't unmarshal request body: %v", err))
		return
	}

	token, _ := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return ctrl.jwtSecretKey, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		httptools.ErrResponse(w, http.StatusInternalServerError, errors.New("unable to retrieve claims from JWT token"))
		return
	}

	username, ok := claims["username"].(string)
	if !ok {
		httptools.ErrResponse(w, http.StatusInternalServerError, errors.New("unable to retrieve username from JWT token"))
		return
	}

	result := ctrl.refreshTokensCollection.FindOne(context.Background(), bson.D{{"username", username}, {"token", req.RefreshToken}})
	if result.Err() != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("refresh token not found: %v", result.Err()))
		return
	}
	var refreshTokensFromDb refreshTokens
	if err := result.Decode(&refreshTokensFromDb); err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't decode refresh token: %v", err))
		return
	}
	if refreshTokensFromDb.ExpireAt.Before(time.Now()) {
		httptools.ErrResponse(w, http.StatusUnauthorized, errors.New("refresh token expired"))
		return
	}
	payload := jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * time.Duration(ctrl.config.AccessTokenExpiredHours)).Unix(),
		"username": username,
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	newAccessToken, err := newToken.SignedString(ctrl.jwtSecretKey)
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't sign token: %v", err))
		return
	}
	payload = jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * time.Duration(ctrl.config.RefreshTokenExpiredHours)).Unix(),
		"username": username,
	}
	newRefreshToken, err := ctrl.signedJwtToken(payload)
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't sign refresh token: %v", err))
		return
	}
	upsert := true
	_, err = ctrl.refreshTokensCollection.ReplaceOne(context.Background(), bson.D{{"username", username}}, refreshTokens{Username: username, Token: newRefreshToken, ExpireAt: time.Now().Add(time.Hour * 24 * 7)}, &options.ReplaceOptions{Upsert: &upsert})
	if err != nil {
		httptools.ErrResponse(w, http.StatusInternalServerError, fmt.Errorf("can't insert refresh token: %v", err))
		return
	}

	httptools.SuccessResponse(w, LoginResponse{AccessToken: newAccessToken, RefreshToken: newRefreshToken})
}
