package server

import (
	"context"
	"fmt"
	"github.com/VxVxN/my_finances/internal/config"
	"github.com/VxVxN/my_finances/internal/controllers"
	"github.com/VxVxN/my_finances/internal/controllers/order"
	"github.com/VxVxN/my_finances/pkg/httptools"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"time"
)

type Server struct {
	OrderController  *order.Controller
	CommonController *controllers.Controller
	client           *mongo.Client
	cfg              *config.Config
}

func Init() (*Server, error) {
	cfg, err := config.Init("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("can't init config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoUrl))
	if err != nil {
		return nil, fmt.Errorf("can't connect to mongodb: %v", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("can't ping to mongodb, %w", err)
	}

	commonController, err := controllers.Init(client)
	if err != nil {
		return nil, fmt.Errorf("can't init common controller: %v", err)
	}

	return &Server{
		OrderController:  order.Init(client),
		CommonController: commonController,
		client:           client,
		cfg:              cfg,
	}, nil
}

func (server *Server) Stop() {
	if err := server.client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

func (server *Server) ListenAndServe(handler http.Handler) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", server.cfg.Port), handler)
}

func (server *Server) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !server.cfg.DisableAuth {
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
		}
		next.ServeHTTP(w, r)
	}
}
