package card

import (
	"github.com/dmytro-kucherenko/card-service/internal/pkg/errors"
)

const errRange = 1 * errors.Range

const (
	ErrNumberInvalid errors.ErrCode = iota + errRange
	ErrNumberCheckDigitInvalid
	ErrCardExpired
)

func IsError(err error) bool {
	return errors.IsRange(err, errRange)
}
