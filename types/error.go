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
	// Indicates that the error is not specific
	ERR_GENERIC ErrorCode = 0

	// Indicates that request body cannot be read.
	ERR_BODY_UNREADABLE ErrorCode = 1

	// Indicates that byte slice cannot be unmarshaled.
	ERR_UNMARSHAL ErrorCode = 2

	// Indicates that something went wrong while inserting an entry to database.
	ERR_DB_INSERT ErrorCode = 3

	// Indicates that something went wrong while listing entries from database.
	ERR_DB_LIST ErrorCode = 4

	// Indicates that an error occured while decoding db results.
	ERR_DB_DECODE ErrorCode = 5

	// Indicates that an error occured while deleting a db entry.
	ERR_DB_DELETE ErrorCode = 6

	// Indicates that the specified entity is not found.
	ERR_NOT_FOUND ErrorCode = 7

	// Indicates that an error occured while updating an entity.
	ERR_DB_UPDATE ErrorCode = 8

	// Indicates that no valid field has been specified for an entity update.
	ERR_NOTHING_TO_UPDATE ErrorCode = 9

	// Indicates that DbURL field of the config has been left empty.
	ERR_CFG_DB_URL_NOT_SPECIFIED ErrorCode = 10

	// Indicates that DbName field of the config has been left empty.
	ERR_CFG_DB_NAME_NOT_SPECIFIED ErrorCode = 11

	// Indicates that http request has wrong http method.
	ERR_WRONG_HTTP_METHOD ErrorCode = 12

	// Indicates that specified category id when creating a menu item does not exist.
	ERR_DB_CATEGORY_NOT_FOUND ErrorCode = 13

	// Indicates a wrong value has been specified for logrus formatter in config.
	ERR_CFG_LOGRUS_FORMATTER_INVALID ErrorCode = 14

	// Indicates an error occured when gettin an entry from db.
	ERR_DB_GET ErrorCode = 15
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
func Errf(code ErrorCode, format string, args ...interface{}) error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
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
