package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/umutozd/restaurant-backend/types"
)

func (s *server) writeResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	if data == nil {
		data = map[string]string{
			"status": "ok",
		}
	}
	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(
			fmt.Sprintf("error marshaling server response: %v", err),
		))
		return
	}

	w.WriteHeader(statusCode)
	_, _ = w.Write(b)
}

func (s *server) writeError(w http.ResponseWriter, logger *logrus.Entry, err error) {
	if logger != nil {
		logger.Error(err)
	}

	switch val := err.(type) {
	case *types.Error:
		realErr := val
		w.WriteHeader(realErr.Code.ToHTTPStatus())
		w.Write([]byte(realErr.Error()))

	default:
		// convert the given error to an instance of types.Error and write it
		logrus.Warnf("WriteError: specified argument's type is not *types.Error, but: %T", err)
		realErr := &types.Error{
			Code:    types.ERR_GENERIC,
			Message: err.Error(),
		}
		w.WriteHeader(realErr.Code.ToHTTPStatus())
		w.Write([]byte(realErr.Error()))
	}
}
