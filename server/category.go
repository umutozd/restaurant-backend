package server

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/umutozd/restaurant-backend/types"
	"github.com/umutozd/restaurant-backend/types/requests"
)

func (s *server) CreateCategory(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("endpoint", "CreateCategory")

	// read request body
	category := &types.Category{}
	if err := category.UnmarshalBody(r.Body); err != nil {
		s.writeError(w, logger, err)
		return
	}
	logger = logger.WithField("category", category)
	logger.Info()

	created, err := s.storage.CreateCategory(r.Context(), category)
	if err != nil {
		s.writeError(w, logger, err)
		return
	}

	s.writeResponse(w, created, http.StatusOK)
}

func (s *server) ListCategories(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("endpoint", "ListCategories")
	logger.Info()

	categories, err := s.storage.ListCategories(r.Context())
	if err != nil {
		s.writeError(w, logger, err)
		return
	}

	s.writeResponse(w, categories, http.StatusOK)
}

func (s *server) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("endpoint", "UpdateCategory")

	// read request body
	req := &requests.UpdateCategoryReq{}
	if err := req.UnmarshalBody(r.Body); err != nil {
		s.writeError(w, logger, err)
		return
	}
	logger = logger.WithField("id", req.Category.ID)
	logger.Info()

	updated, err := s.storage.UpdateCategory(r.Context(), req)
	if err != nil {
		s.writeError(w, logger, err)
		return
	}

	s.writeResponse(w, updated, http.StatusOK)
}

func (s *server) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("endpoint", "DeleteCategory")
	logger.Info()

	// read request body
	req := &requests.DeleteCategoryReq{}
	if err := req.UnmarshalBody(r.Body); err != nil {
		s.writeError(w, logger, err)
		return
	}
	logger = logger.WithField("id", req.ID)
	logger.Info()

	if err := s.storage.DeleteCategory(r.Context(), req); err != nil {
		s.writeError(w, logger, err)
		return
	}

	s.writeResponse(w, nil, http.StatusOK)
}
