package card

import "github.com/dmytro-kucherenko/card-service/internal/api/rest/pkg/models"

type ValidateRequest struct {
	Number string `json:"number" validate:"required"`
	Month  uint8  `json:"month" validate:"required,min=1,max=12"`
	Year   uint16 `json:"year" validate:"required,min=1970,max=65535"`
} // @name CardValidateRequest

type ValidateResponse models.Response // @name CardValidateResponse
