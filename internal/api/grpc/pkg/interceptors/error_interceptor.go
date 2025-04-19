package interceptors

import (
	"context"

	grpcErrors "github.com/dmytro-kucherenko/card-service/internal/api/grpc/pkg/errors"
	appErrors "github.com/dmytro-kucherenko/card-service/internal/pkg/errors"
	"github.com/dmytro-kucherenko/card-service/internal/pkg/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	statusErrors "google.golang.org/grpc/status"
)

func getError(code appErrors.ErrCode, status codes.Code, message string) error {
	satusErr, _ := statusErrors.New(status, message).WithDetails(&errdetails.ErrorInfo{
		Reason: code.String(),
	})

	return satusErr.Err()
}

func ErrorUnary(
	logger log.Logger,
) func(context.Context, any, *grpc.UnaryServerInfo, grpc.UnaryHandler) (any, error) {
	return func(ctx context.Context, request any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (result any, err error) {
		defer func() {
			if errPanic := recover(); errPanic != nil {
				logger.Error("Panic recovery: ", errPanic)
				err = getError(appErrors.ErrInternal, codes.Internal, "internal server error")
			}
		}()

		result, err = handler(ctx, request)
		if err == nil {
			return
		}

		code, okCode := appErrors.Code(err)
		status, okStatus := grpcErrors.Status(err)

		if !okCode || !okStatus {
			logger.Error(err.Error())

			return nil, getError(appErrors.ErrInternal, codes.Internal, "internal server error")
		}

		return nil, getError(code, status, err.Error())
	}
}
