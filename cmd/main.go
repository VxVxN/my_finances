package main

import (
	financeserver "github.com/VxVxN/my_finances/internal/server"
	"log"
)

func main() {
	server := financeserver.Init()
	server.Handle("POST /order/create", server.Controller.CreateOrder)
	server.Handle("POST /order/remove", server.Controller.RemoveOrder)
	server.Handle("GET /orders", server.Controller.Orders)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Cannot listen server: %v", err)
	}
}
