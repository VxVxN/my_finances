package main

import (
	financeserver "github.com/VxVxN/my_finances/internal/server"
	"log"
)

func main() {
	server, err := financeserver.Init()
	if err != nil {
		log.Fatalf("Error initializing server: %v", err)
	}
	defer server.Stop()
	server.Handle("GET /", server.CommonController.Index)

	server.Handle("POST /order/create", server.OrderController.CreateOrder)
	server.Handle("POST /order/remove", server.OrderController.RemoveOrder)
	server.Handle("GET /orders", server.OrderController.Orders)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Cannot listen server: %v", err)
	}
}
