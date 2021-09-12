package server

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/umutozd/restaurant-backend/types/requests"
)

func (s *server) UpdateCart(w http.ResponseWriter, r *http.Request) {
	logger := s.getLoggerFromRequest(r)

	req := &requests.UpdateCartReq{}
	if err := req.UnmarshalBody(r.Body); err != nil {
		s.writeError(w, logger, err)
		return
	}
	logger = logger.WithFields(logrus.Fields{
		"cart_id": req.ID,
		"add":     req.Add,
		"remove":  req.Remove,
	})
	logger.Info()

	cart, err := s.storage.UpdateCart(r.Context(), req)
	if err != nil {
		s.writeError(w, logger, err)
		return
	}

	s.writeResponse(w, cart, http.StatusOK)
}
