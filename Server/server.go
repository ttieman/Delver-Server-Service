package Server

import (
	"github.com/gorilla/mux" // Package for routing and dispatching HTTP requests
)

type Server struct {
	*mux.Router //embed the mux.Router to handle HTTP requests
}

func NewServer() *Server {

	server := &Server{
		Router: mux.NewRouter(),
		users:  []User{},
	}
	return server
}
