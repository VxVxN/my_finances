package server

import (
	"context"
	"fmt"
	"github.com/VxVxN/my_finances/internal/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"time"
)

type Server struct {
	Controller *controllers.Controller
	client     *mongo.Client
}

func Init() (*Server, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		return nil, fmt.Errorf("can't connect to mongodb: %v", err)
	}
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("can't ping to mongodb, %w", err)
	}

	return &Server{
		Controller: controllers.Init(client),
		client:     client,
	}, nil
}

func (server *Server) Stop() {
	if err := server.client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

func (server *Server) ListenAndServe() error {
	return http.ListenAndServe(":8080", nil) // todo address move to config
}

func (server *Server) Handle(route string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(route, handler)
}
