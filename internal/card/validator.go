package card

import (
	"regexp"
	"strconv"
	"time"

	"github.com/dmytro-kucherenko/card-service/internal/pkg/errors"
)

const (
	pattern         string = "^[0-9]{8,19}$"
	luhnCoef        int    = 2
	luhnMaxDigitSum int    = 9
	luhnModulo      int    = 10
)

func validateNumberPattern(number string) error {
	ok := regexp.MustCompile(pattern).MatchString(number)
	if !ok {
		return errors.NewAppError(ErrNumberInvalid, "card number is not numeric or has invalid length")
	}

	return nil
}

func validateNumberLuhn(number string) error {
	sum := 0
	for i := len(number) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return errors.NewAppError(ErrNumberInvalid, "card number is not numeric")
		}

		isSecondDigit := (len(number)-i)%2 == 0
		if isSecondDigit {
			digit *= luhnCoef
			if digit > luhnMaxDigitSum {
				digit -= luhnMaxDigitSum
			}
		}

		sum += digit
	}

	valid := sum%luhnModulo == 0
	if !valid {
		return errors.NewAppError(ErrNumberCheckDigitInvalid, "card number check digit is invalid")
	}

	return nil
}

func validateNumber(number string) error {
	validations := []func(string) error{
		validateNumberPattern,
		validateNumberLuhn,
	}

	for _, validate := range validations {
		err := validate(number)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateExpirationDate(month uint8, year uint16) error {
	startOfMonth := time.Date(int(year), time.Month(month), 1, 23, 59, 59, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	now := time.Now()
	if endOfMonth.Before(now) {
		return errors.NewAppError(ErrCardExpired, "card is expired")
	}

	return nil
}

func Validate(card Item) error {
	err := validateNumber(card.Number)
	if err != nil {
		return err
	}

	return validateExpirationDate(card.Month, card.Year)
}
