package errors

import (
	"fmt"

	"github.com/dmytro-kucherenko/card-service/internal/pkg/utils"
)

type ErrCode uint16

const Range = 100

const (
	ErrInternal ErrCode = iota + 1
	ErrValidation
)

func (code ErrCode) String() string {
	return utils.PadLeft(fmt.Sprint(uint16(code)), '0', 3)
}
