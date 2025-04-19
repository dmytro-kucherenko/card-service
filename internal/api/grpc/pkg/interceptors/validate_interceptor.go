package interceptors

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	grpcErrors "github.com/dmytro-kucherenko/card-service/internal/api/grpc/pkg/errors"
	appErrors "github.com/dmytro-kucherenko/card-service/internal/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func ValidateUnary() func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error) {
	validator, _ := protovalidate.New()

	return func(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		message := request.(proto.Message)
		if err := validator.Validate(message); err != nil {
			return nil, grpcErrors.NewGRPCError(codes.InvalidArgument, appErrors.ErrValidation, err.Error())
		}

		return handler(ctx, request)
	}
}
