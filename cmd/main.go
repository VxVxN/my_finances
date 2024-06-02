package main

import (
	"fmt"
	"github.com/VxVxN/my_finances/internal/controllers"
	financeserver "github.com/VxVxN/my_finances/internal/server"
	"github.com/VxVxN/my_finances/pkg/httptools"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
)

func main() {
	server, err := financeserver.Init()
	if err != nil {
		log.Fatalf("Error initializing server: %v", err)
	}
	defer server.Stop()

	router := http.NewServeMux()
	router.HandleFunc("GET /", server.CommonController.Index)
	router.HandleFunc("POST /register", server.CommonController.Register)
	router.HandleFunc("POST /login", server.CommonController.Login)

	router.HandleFunc("POST /order/create", AuthMiddleware(server.OrderController.CreateOrder))
	router.HandleFunc("POST /order/remove", AuthMiddleware(server.OrderController.RemoveOrder))
	router.HandleFunc("GET /orders", AuthMiddleware(server.OrderController.Orders))

	if err := server.ListenAndServe(router); err != nil {
		log.Fatalf("Cannot listen server: %v", err)
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := r.Cookie("access_token")
		if err != nil {
			httptools.ErrResponse(w, http.StatusBadRequest, fmt.Errorf("cannot get access token: %v", err))
			return
		}
		token, err := jwt.Parse(accessToken.Value, func(token *jwt.Token) (interface{}, error) {
			return controllers.JwtSecretKey, nil
		})

		if err != nil {
			httptools.ErrResponse(w, http.StatusUnauthorized, fmt.Errorf("failed to parse access token: %v", err))
			return
		}

		if !token.Valid {
			httptools.ErrResponse(w, http.StatusUnauthorized, fmt.Errorf("invalid access token: %v", err))
			return
		}
		next.ServeHTTP(w, r)
	}
}
