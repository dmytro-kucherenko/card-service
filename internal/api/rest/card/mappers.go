package card

import (
	"net/http"

	restErrors "github.com/dmytro-kucherenko/card-service/internal/api/rest/pkg/errors"
	"github.com/dmytro-kucherenko/card-service/internal/card"
	appErrors "github.com/dmytro-kucherenko/card-service/internal/pkg/errors"
)

func mapValidateRequest(request ValidateRequest) card.Item {
	return card.Item(request)
}

func mapError(err error) error {
	if !card.IsError(err) {
		return err
	}

	status := http.StatusInternalServerError
	switch code, _ := appErrors.Code(err); code {
	case card.ErrNumberInvalid, card.ErrNumberCheckDigitInvalid, card.ErrCardExpired:
		status = http.StatusBadRequest
	}

	return restErrors.NewHTTPFromError(status, err)
}
