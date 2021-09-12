package server

import (
	"net/http"

	"github.com/gorilla/mux"
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
	store, err := storage.NewStorage(cfg.DbURL, cfg.DbName, cfg.DBCronInterval)
	if err != nil {
		return nil, err
	}
	return &server{
		cfg:     cfg,
		storage: store,
	}, nil
}

func (s *server) Listen() error {
	r := mux.NewRouter()
	r.HandleFunc(EndpointListMenu.String(), s.getHandler(EndpointListMenu))

	r.HandleFunc(EndpointCreateMenuItem.String(), s.getHandler(EndpointCreateMenuItem))
	r.HandleFunc(EndpointListMenuItems.String(), s.getHandler(EndpointListMenuItems))
	r.HandleFunc(EndpointUpdateMenuItem.String(), s.getHandler(EndpointUpdateMenuItem))
	r.HandleFunc(EndpointDeleteMenuItem.String(), s.getHandler(EndpointDeleteMenuItem))

	r.HandleFunc(EndpointCreateCategory.String(), s.getHandler(EndpointCreateCategory))
	r.HandleFunc(EndpointListCategories.String(), s.getHandler(EndpointListCategories))
	r.HandleFunc(EndpointUpdateCategory.String(), s.getHandler(EndpointUpdateCategory))
	r.HandleFunc(EndpointDeleteCategory.String(), s.getHandler(EndpointDeleteCategory))

	r.HandleFunc(EndpointUpdateCart.String(), s.getHandler(EndpointUpdateCart))
	r.HandleFunc(EndpointUpdateCartItemsStatus.String(), s.getHandler(EndpointUpdateCartItemsStatus))
	http.Handle("/", r)

	if err := http.ListenAndServe(s.cfg.GetPort(), r); err != nil {
		return err
	}

	return nil
}
