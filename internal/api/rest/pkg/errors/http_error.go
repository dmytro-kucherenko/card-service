package errors

import (
	"errors"
	"net/http"

	appErrors "github.com/dmytro-kucherenko/card-service/internal/pkg/errors"
)

type HTTPError struct {
	error
	status int
}

func NewHTTPError(status int, code appErrors.ErrCode, message string) error {
	return &HTTPError{
		error:  appErrors.NewAppError(code, message),
		status: status,
	}
}

func NewHTTPFromError(status int, err error) error {
	return &HTTPError{
		error:  err,
		status: status,
	}
}

func (e *HTTPError) Unwrap() error {
	return e.error
}

func Status(err error) (int, bool) {
	var httpErr *HTTPError
	ok := errors.As(err, &httpErr)
	if !ok {
		return http.StatusInternalServerError, false
	}

	return httpErr.status, true
}
