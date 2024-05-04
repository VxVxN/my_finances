package server

import (
	"github.com/VxVxN/my_finances/internal/controllers"
	"net/http"
)

type Server struct {
	Controller *controllers.Controller
}

func Init() *Server {
	return &Server{
		Controller: controllers.Init(),
	}
}

func (server *Server) ListenAndServe() error {
	return http.ListenAndServe(":8080", nil) // todo address move to config
}

func (server *Server) Handle(route string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(route, handler)
}
