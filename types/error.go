package types

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// ErrorCode specifies what kind of error the Error type contains.
type ErrorCode int32

const (
	// Indicates that request body cannot be read.
	ERR_BODY_UNREADABLE ErrorCode = 0

	// Indicates that byte slice cannot be unmarshaled.
	ERR_UNMARSHAL ErrorCode = 1
)

func (ec ErrorCode) ToHTTPStatus() int {
	switch ec {
	case 0, 1:
		return http.StatusInternalServerError
	default:
		return http.StatusOK
	}
}

// Error is the custom error type that is used accross the project
// which also implements the error interface.
type Error struct {
	Code    ErrorCode `json:"code,omitempty"`
	Message string    `json:"message,omitempty"`
}

// Error is the method for implementing error interface. It returns
// the json string of the Error object.
func (e *Error) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		logrus.WithError(err).Error("error marshaling Error struct")
		return fmt.Sprintf(`{"message":%q, "code":%d}`, e.Message, e.Code)
	}
	return string(b)
}

// ErrByString returns an Error with the specified code and message.
func ErrByString(code ErrorCode, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// ErrByString returns an Error with the specified code and message as
// the text within err.
func Err(code ErrorCode, err error) error {
	return &Error{
		Code:    code,
		Message: err.Error(),
	}
}
