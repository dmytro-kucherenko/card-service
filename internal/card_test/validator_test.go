package card_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/dmytro-kucherenko/card-service/internal/card"
	"github.com/dmytro-kucherenko/card-service/internal/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const validNumber = "4111111111111111"

func getFromNow(addMonths int) (month uint8, year uint16) {
	now := time.Now().AddDate(0, addMonths, 0)
	month = uint8(now.Month())
	year = uint16(now.Year())

	return
}

func ExampleValidate() {
	month, year := getFromNow(0)
	item := card.Item{Number: validNumber, Month: month, Year: year}

	err := card.Validate(item)

	fmt.Println("Valid:", err == nil)
	// Output:
	// Valid: true
}

func TestValidate(t *testing.T) {
	month, year := getFromNow(0)
	expMonth, expYear := getFromNow(-1)

	tests := []struct {
		name   string
		number string
		month  uint8
		year   uint16
		ok     bool
		code   errors.ErrCode
	}{
		{
			name:   "Valid",
			number: validNumber,
			month:  month,
			year:   year,
			ok:     true,
		},
		{
			name:   "InvalidShort",
			number: "4111118",
			month:  month,
			year:   year,
			code:   card.ErrNumberInvalid,
		},
		{
			name:   "InvalidLong",
			number: "41111111111111111115",
			month:  month,
			year:   year,
			code:   card.ErrNumberInvalid,
		},
		{
			name:   "InvalidNotNumeric",
			number: "41111111a_111111",
			month:  month,
			year:   year,
			code:   card.ErrNumberInvalid,
		},
		{
			name:   "InvalidCheckDigit",
			number: "4111111111111112",
			month:  month,
			year:   year,
			code:   card.ErrNumberCheckDigitInvalid,
		},
		{
			name:   "InvalidExpired",
			number: validNumber,
			month:  expMonth,
			year:   expYear,
			code:   card.ErrCardExpired,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := card.Validate(card.Item{Number: test.number, Month: test.month, Year: test.year})

			if test.ok {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				require.IsType(t, &errors.AppError{}, err)

				code, _ := errors.Code(err)
				assert.Equal(t, test.code, code)
			}
		})
	}
}

func BenchmarkValidate(b *testing.B) {
	month, year := getFromNow(0)
	item := card.Item{Number: validNumber, Month: month, Year: year}

	b.ResetTimer()

	for b.Loop() {
		card.Validate(item)
	}
}
