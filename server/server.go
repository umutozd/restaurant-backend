package server

import (
	"net/http"
)

type Server interface {
	Listen() error
}

type server struct {
	cfg *Config
}

func NewServer(cfg *Config) Server {
	return &server{
		cfg: cfg,
	}
}

func (s *server) Listen() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/menu/list", nil)

	mux.HandleFunc("/api/v1/menu/items/list", nil)
	mux.HandleFunc("/api/v1/menu/items/add", nil)
	mux.HandleFunc("/api/v1/menu/items/update", nil)
	mux.HandleFunc("/api/v1/menu/items/delete", nil)

	mux.HandleFunc("/api/v1/category/list", nil)
	mux.HandleFunc("/api/v1/category/add", nil)
	mux.HandleFunc("/api/v1/category/update", nil)
	mux.HandleFunc("/api/v1/category/delete", nil)
	return nil
}
