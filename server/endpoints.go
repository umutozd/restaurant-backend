package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/umutozd/restaurant-backend/types"
)

type loggerKey struct{}

type endpoint string

const (
	EndpointListMenu       endpoint = "/api/v1/menu/list"
	EndpointCreateMenuItem endpoint = "/api/v1/menu/items/create"
	EndpointListMenuItems  endpoint = "/api/v1/menu/items/list"
	EndpointUpdateMenuItem endpoint = "/api/v1/menu/items/update"
	EndpointDeleteMenuItem endpoint = "/api/v1/menu/items/delete"

	EndpointCreateCategory endpoint = "/api/v1/category/create"
	EndpointListCategories endpoint = "/api/v1/category/list"
	EndpointUpdateCategory endpoint = "/api/v1/category/update"
	EndpointDeleteCategory endpoint = "/api/v1/category/delete"

	EndpointUpdateCart            endpoint = "/api/v1/cart/update"
	EndpointUpdateCartItemsStatus endpoint = "/api/v1/cart/update/items/status"
)

var allEndpoints = []endpoint{
	EndpointListMenu, EndpointCreateMenuItem, EndpointListMenuItems, EndpointUpdateMenuItem,
	EndpointDeleteMenuItem, EndpointCreateCategory, EndpointListCategories,
	EndpointUpdateCategory, EndpointDeleteCategory, EndpointUpdateCart, EndpointUpdateCartItemsStatus,
}

func (e endpoint) String() string {
	return string(e)
}

func (s *server) getLoggerFromRequest(r *http.Request) *logrus.Entry {
	val := r.Context().Value(loggerKey{})
	switch tt := val.(type) {
	case *logrus.Entry:
		return tt
	default:
		path := r.URL.Path
		for _, e := range allEndpoints {
			if e.String() == path {
				return logrus.WithField("endpoint", e)
			}
		}
		return logrus.WithField("endpoint", path)
	}
}

// getHandler maps the giben endpoint to its handlerFunc. The returned
// function will be the wrapped version of server methods. Each of those
// methods will be able to retrieve their own logger objects from context.
func (s *server) getHandler(e endpoint) http.HandlerFunc {
	var h http.HandlerFunc
	var method string
	switch e {
	case EndpointListMenu:
		h = s.ListMenu
		method = http.MethodGet
	case EndpointCreateMenuItem:
		h = s.CreateMenuItem
		method = http.MethodPost
	case EndpointListMenuItems:
		h = s.ListMenuItems
		method = http.MethodGet
	case EndpointUpdateMenuItem:
		h = s.UpdateMenuItem
		method = http.MethodPost
	case EndpointDeleteMenuItem:
		h = s.DeleteMenuItem
		method = http.MethodDelete
	case EndpointCreateCategory:
		h = s.CreateCategory
		method = http.MethodPost
	case EndpointListCategories:
		h = s.ListCategories
		method = http.MethodGet
	case EndpointUpdateCategory:
		h = s.UpdateCategory
		method = http.MethodPost
	case EndpointDeleteCategory:
		h = s.DeleteCategory
		method = http.MethodDelete
	case EndpointUpdateCart:
		h = s.UpdateCart
		method = http.MethodPost
	case EndpointUpdateCartItemsStatus:
		h = s.UpdateCartItemsStatus
		method = http.MethodPost
	default:
		return func(w http.ResponseWriter, r *http.Request) {
			logger := s.getLoggerFromRequest(r)
			msg := fmt.Sprintf("requested path %q is not found", r.URL.Path)
			logger.Warn(msg)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(msg))
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		logger := logrus.WithField("endpoint", e)
		if r.Method != method {
			s.writeError(w, logger, types.Errf(types.ERR_WRONG_HTTP_METHOD, "for endpoint %q, http method must be %s", e, method))
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), loggerKey{}, logger))
		h(w, r)
	}
}
