package server

import (
	"net/http"

	"github.com/umutozd/restaurant-backend/storage"
)

type Server interface {
	Listen() error
}

type server struct {
	cfg     *Config
	storage storage.Storage
}

func NewServer(cfg *Config) (Server, error) {
	store, err := storage.NewStorage(cfg.DbURL, cfg.DbName)
	if err != nil {
		return nil, err
	}
	return &server{
		cfg:     cfg,
		storage: store,
	}, nil
}

func (s *server) Listen() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/menu/list", s.ListMenu)

	mux.HandleFunc("/api/v1/menu/items/create", s.CreateMenuItem)
	mux.HandleFunc("/api/v1/menu/items/list", s.ListMenuItems)
	mux.HandleFunc("/api/v1/menu/items/update", s.UpdateMenuItem)
	mux.HandleFunc("/api/v1/menu/items/delete", s.DeleteMenuItem)

	mux.HandleFunc("/api/v1/category/create", s.CreateCategory)
	mux.HandleFunc("/api/v1/category/list", s.ListCategories)
	mux.HandleFunc("/api/v1/category/update", s.UpdateCategory)
	mux.HandleFunc("/api/v1/category/delete", s.DeleteCategory)

	if err := http.ListenAndServe(s.cfg.GetPort(), mux); err != nil {
		return err
	}

	return nil
}
