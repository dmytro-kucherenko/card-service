package card

import (
	"context"

	gen "github.com/dmytro-kucherenko/card-service/api/gen/grpc/card"
	"github.com/dmytro-kucherenko/card-service/internal/card"
	"google.golang.org/grpc"
)

type Handler struct {
	gen.UnimplementedServiceServer
}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) Init(server *grpc.Server) {
	gen.RegisterServiceServer(server, handler)
}

func (handler *Handler) Validate(ctx context.Context, request *gen.ValidateRequest) (*gen.ValidateResponse, error) {
	err := card.Validate(mapValidateRequest(request))
	if err != nil {
		return nil, mapError(err)
	}

	return &gen.ValidateResponse{Valid: true}, nil
}
