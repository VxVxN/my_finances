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
	server.Handle("GET /", server.Controller.Index)
	server.Handle("POST /order/create", server.Controller.CreateOrder)
	server.Handle("POST /order/remove", server.Controller.RemoveOrder)
	server.Handle("GET /orders", server.Controller.Orders)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Cannot listen server: %v", err)
	}
}
