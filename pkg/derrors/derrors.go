package derrors

import (
	"errors"
	"net/http"

	"github.com/mehrdadjalili/facegram_common/pkg/translator/messages"
)

type (
	status uint

	serverError struct {
		status  status
		message string
	}
)

const (
	_ status = iota
	StatusNoContent
	StatusBadRequest
	StatusUnauthorized
	StatusPaymentRequired
	StatusForbidden
	StatusNotFound
	StatusNotAcceptable
	StatusUnsupportedMediaType
	StatusLocked
	StatusUpgradeRequired
	StatusTooManyRequests
	StatusInternalServerError
	StatusMethodNotAllowed
	StatusUnexpected
)

var (
	httpErrors = map[status]int{
		StatusNoContent:            http.StatusNoContent,
		StatusBadRequest:           http.StatusBadRequest,
		StatusUnauthorized:         http.StatusUnauthorized,
		StatusPaymentRequired:      http.StatusPaymentRequired,
		StatusForbidden:            http.StatusForbidden,
		StatusNotFound:             http.StatusNotFound,
		StatusNotAcceptable:        http.StatusNotAcceptable,
		StatusUnsupportedMediaType: http.StatusUnsupportedMediaType,
		StatusLocked:               http.StatusLocked,
		StatusUpgradeRequired:      http.StatusUpgradeRequired,
		StatusTooManyRequests:      http.StatusTooManyRequests,
		StatusInternalServerError:  http.StatusInternalServerError,
		StatusMethodNotAllowed:     http.StatusMethodNotAllowed,
		StatusUnexpected:           http.StatusInternalServerError,
	}
)

//New is constructor of the errors package
func New(status status, msg string) error {
	return serverError{
		status:  status,
		message: msg,
	}
}

//Error return message of error
func (e serverError) Error() string {
	return e.message
}

//HttpError convert kind of error to Http status error
//if error type is not serverError return 400 status code
func HttpError(err error) (string, int) {
	var serverErr serverError

	ok := errors.As(err, &serverErr)
	if !ok {
		return messages.InternalServerError, http.StatusInternalServerError
	}

	code, ok := httpErrors[serverErr.status]
	if !ok {
		return serverErr.message, http.StatusBadRequest
	}

	return serverErr.message, code

}

func As(err error) bool {
	return errors.As(err, &serverError{})
}
