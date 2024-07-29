package main

import (
	"log"
	"net/http"

	financeserver "github.com/VxVxN/my_finances/internal/server"
	"github.com/VxVxN/my_finances/pkg/httptools"
	"github.com/rs/cors"
)

// Commit is a git commit hash, set by ldflags
var Commit string

func main() {
	server, err := financeserver.Init(Commit)
	if err != nil {
		log.Fatalf("Error initializing server: %v", err)
	}
	if err = server.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer server.Stop()

	commonMiddleware := []httptools.Middleware{
		server.AuthMiddleware,
		server.LogsMiddleware,
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /", server.LogsMiddleware(server.CommonController.Index))
	router.HandleFunc("POST /register", server.LogsMiddleware(server.CommonController.Register))
	router.HandleFunc("POST /login", server.LogsMiddleware(server.CommonController.Login))
	router.HandleFunc("POST /refresh-token", server.LogsMiddleware(server.CommonController.RefreshToken))

	router.HandleFunc("POST /order/create", httptools.MultipleMiddleware(server.OrderController.CreateOrder, commonMiddleware...))
	router.HandleFunc("POST /order/remove", httptools.MultipleMiddleware(server.OrderController.RemoveOrder, commonMiddleware...))
	router.HandleFunc("GET /orders", httptools.MultipleMiddleware(server.OrderController.Orders, commonMiddleware...))
	router.HandleFunc("GET /balance", httptools.MultipleMiddleware(server.OrderController.Balance, commonMiddleware...))

	router.HandleFunc("POST /chart/historical-balance", httptools.MultipleMiddleware(server.ChartController.HistoricalBalance, commonMiddleware...))

	routerWithCors := cors.Default().Handler(router)
	if err := server.ListenAndServe(routerWithCors); err != nil {
		log.Fatalf("Cannot listen server: %v", err)
	}
}
