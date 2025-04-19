package errors

import (
	"errors"

	appErrors "github.com/dmytro-kucherenko/card-service/internal/pkg/errors"
	"google.golang.org/grpc/codes"
)

type GRPCError struct {
	error
	status codes.Code
}

func NewGRPCError(status codes.Code, code appErrors.ErrCode, message string) error {
	return &GRPCError{
		error:  appErrors.NewAppError(code, message),
		status: status,
	}
}

func NewGRPCFromError(status codes.Code, err error) error {
	return &GRPCError{
		error:  err,
		status: status,
	}
}

func (e *GRPCError) Unwrap() error {
	return e.error
}

func Status(err error) (codes.Code, bool) {
	var grpcErr *GRPCError
	ok := errors.As(err, &grpcErr)
	if !ok {
		return codes.Internal, false
	}

	return grpcErr.status, true
}
