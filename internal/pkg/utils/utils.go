package utils

import (
	"strings"
)

func PadLeft(text string, char rune, length int) string {
	padLength := length - len(text)
	if padLength <= 0 {
		return text
	}

	return strings.Repeat(string(char), padLength) + text
}
