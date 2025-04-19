package errors

import (
	"errors"
)

type AppError struct {
	error
	code ErrCode
}

func NewAppError(code ErrCode, message string) error {
	return &AppError{
		error: errors.New(message),
		code:  code,
	}
}

func Code(err error) (ErrCode, bool) {
	var appErr *AppError
	ok := errors.As(err, &appErr)
	if !ok {
		return ErrInternal, false
	}

	return appErr.code, true
}

func IsRange(err error, errRange int) bool {
	code, ok := Code(err)
	isRange := (int(code)-errRange)/Range == 0

	return ok && isRange
}
