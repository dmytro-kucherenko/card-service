package card

import (
	"context"
	"net/http"

	"github.com/dmytro-kucherenko/card-service/internal/api/rest/pkg/interceptors"
	"github.com/dmytro-kucherenko/card-service/internal/api/rest/pkg/register"
	"github.com/dmytro-kucherenko/card-service/internal/card"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) Init(router register.Router) {
	group := router.Group("/card").Use(interceptors.Map(mapError))
	register.Post(group, "/validate", http.StatusOK, handler.Validate)
}

// @Summary Validate card
// @Tags		Card
// @Accept		json
// @Produce	json
// @Param		body	body		CardValidateRequest	true	"CardData"
// @Success	200		{object}	CardValidateResponse
// @Failure	400		{object}	ErrorResponse
// @Router		/card/validate [post]
func (handler *Handler) Validate(ctx context.Context, request ValidateRequest) (response ValidateResponse, err error) {
	err = card.Validate(mapValidateRequest(request))
	if err != nil {
		return
	}

	return ValidateResponse{Valid: true}, nil
}
