package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// ConvertText function at first determine type of text (Morse code or simple text) and then convert them to each other type.
func ConvertText(text string) string {

	f := func(r rune) bool {
		return r == '-' || r == '.' || r == ' '
	}
	switch strings.ContainsFunc(text, f) {
	case true:
		return morse.ToText(text)
	case false:
		return morse.ToMorse(text)
	default:
		return morse.ErrNoEncoding.Error(morse.ErrNoEncoding{})
	}
}
