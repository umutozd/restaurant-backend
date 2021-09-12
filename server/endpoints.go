package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
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

	EndpointUpdateCart endpoint = "/api/v1/cart/update"
)

var allEndpoints = []endpoint{
	EndpointListMenu, EndpointCreateMenuItem, EndpointListMenuItems, EndpointUpdateMenuItem,
	EndpointDeleteMenuItem, EndpointCreateCategory, EndpointListCategories,
	EndpointUpdateCategory, EndpointDeleteCategory, EndpointUpdateCart,
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
	switch e {
	case EndpointListMenu:
		h = s.ListMenu
	case EndpointCreateMenuItem:
		h = s.CreateMenuItem
	case EndpointListMenuItems:
		h = s.ListMenuItems
	case EndpointUpdateMenuItem:
		h = s.UpdateMenuItem
	case EndpointDeleteMenuItem:
		h = s.DeleteMenuItem
	case EndpointCreateCategory:
		h = s.CreateCategory
	case EndpointListCategories:
		h = s.ListCategories
	case EndpointUpdateCategory:
		h = s.UpdateCategory
	case EndpointDeleteCategory:
		h = s.DeleteCategory
	case EndpointUpdateCart:
		h = s.UpdateCart
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
		r = r.WithContext(context.WithValue(r.Context(), loggerKey{}, logger))
		h(w, r)
	}
}
