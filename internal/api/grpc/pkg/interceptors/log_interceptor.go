package interceptors

import (
	"context"
	"fmt"
	"time"

	"github.com/dmytro-kucherenko/card-service/internal/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func LogUnary(
	logger log.Logger,
) func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		result, err := handler(ctx, request)
		end := time.Since(start)

		latency := float64(end.Microseconds()) / 1000.0
		path := info.FullMethod
		code := status.Code(err)
		message := fmt.Sprintf("[%v] %v %vms", code, path, latency)

		if err == nil {
			logger.Info(message)
		} else {
			logger.Warn(fmt.Sprintf("%v\n%v", message, err.Error()))
		}

		return result, err
	}
}
