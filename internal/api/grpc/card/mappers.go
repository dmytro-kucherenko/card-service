package card

import (
	gen "github.com/dmytro-kucherenko/card-service/api/gen/grpc/card"
	grpcErrors "github.com/dmytro-kucherenko/card-service/internal/api/grpc/pkg/errors"
	"github.com/dmytro-kucherenko/card-service/internal/card"
	appErrors "github.com/dmytro-kucherenko/card-service/internal/pkg/errors"
	"google.golang.org/grpc/codes"
)

func mapValidateRequest(request *gen.ValidateRequest) card.Item {
	return card.Item{
		Number: request.Number,
		Month:  uint8(request.Month),
		Year:   uint16(request.Year),
	}
}

func mapError(err error) error {
	if err == nil {
		return nil
	}

	if !card.IsError(err) {
		return err
	}

	status := codes.Internal
	switch code, _ := appErrors.Code(err); code {
	case card.ErrNumberInvalid, card.ErrNumberCheckDigitInvalid, card.ErrCardExpired:
		status = codes.InvalidArgument
	}

	return grpcErrors.NewGRPCFromError(status, err)
}
