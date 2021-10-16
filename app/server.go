package app

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type server struct {
	router *mux.Router
	db     *gorm.DB
}

func NewServer() *server {
	var s server
	s.router = mux.NewRouter()
	s.database()
	s.routes()

	return &s
}

func (s *server) Serve() error {
    return http.ListenAndServe(":8080", s.router)
}
