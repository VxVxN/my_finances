package main

import (
	financeserver "github.com/VxVxN/my_finances/internal/server"
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
	router.HandleFunc("POST /refresh-token", server.CommonController.RefreshToken)

	router.HandleFunc("POST /order/create", server.AuthMiddleware(server.OrderController.CreateOrder))
	router.HandleFunc("POST /order/remove", server.AuthMiddleware(server.OrderController.RemoveOrder))
	router.HandleFunc("GET /orders", server.AuthMiddleware(server.OrderController.Orders))

	if err := server.ListenAndServe(router); err != nil {
		log.Fatalf("Cannot listen server: %v", err)
	}
}
