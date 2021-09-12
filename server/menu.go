package server

import (
	"net/http"

	"github.com/umutozd/restaurant-backend/types"
	"github.com/umutozd/restaurant-backend/types/requests"
)

func (s *server) ListMenu(w http.ResponseWriter, r *http.Request) {
	logger := s.getLoggerFromRequest(r)
	logger.Info()

	menu, err := s.storage.ListMenu(r.Context())
	if err != nil {
		s.writeError(w, logger, err)
		return
	}

	s.writeResponse(w, menu, http.StatusOK)
}

func (s *server) CreateMenuItem(w http.ResponseWriter, r *http.Request) {
	logger := s.getLoggerFromRequest(r)

	// read request body
	item := &types.MenuItem{}
	if err := item.UnmarshalBody(r.Body); err != nil {
		s.writeError(w, logger, err)
		return
	}
	logger = logger.WithField("item", item)
	logger.Info()

	created, err := s.storage.CreateMenuItem(r.Context(), item)
	if err != nil {
		s.writeError(w, logger, err)
		return
	}

	s.writeResponse(w, created, http.StatusOK)
}

func (s *server) ListMenuItems(w http.ResponseWriter, r *http.Request) {
	logger := s.getLoggerFromRequest(r)
	logger.Info()
	// logger := r.Context().Value(keyLogger{}).(*logrus.Entry)

	items, err := s.storage.ListMenuItems(r.Context())
	if err != nil {
		s.writeError(w, logger, err)
		return
	}
	if items == nil {
		items = []*types.MenuItem{}
	}

	s.writeResponse(w, items, http.StatusOK)
}

func (s *server) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	logger := s.getLoggerFromRequest(r)

	// read request body
	req := &requests.UpdateMenuItemReq{}
	if err := req.UnmarshalBody(r.Body); err != nil {
		s.writeError(w, logger, err)
		return
	}
	logger = logger.WithField("id", req.Item.ID)
	logger.Info()

	updated, err := s.storage.UpdateMenuItem(r.Context(), req)
	if err != nil {
		s.writeError(w, logger, err)
		return
	}

	s.writeResponse(w, updated, http.StatusOK)
}

func (s *server) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	logger := s.getLoggerFromRequest(r)
	logger.Info()

	// read request body
	req := &requests.DeleteMenuItemReq{}
	if err := req.UnmarshalBody(r.Body); err != nil {
		s.writeError(w, logger, err)
		return
	}
	logger = logger.WithField("id", req.ID)
	logger.Info()

	if err := s.storage.DeleteMenuItem(r.Context(), req); err != nil {
		s.writeError(w, logger, err)
		return
	}

	s.writeResponse(w, nil, http.StatusOK)
}
