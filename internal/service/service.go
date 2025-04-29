package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// ConvertText function at first determine type of text (Morse code or simple text) and then convert them to each other type.
func ConvertText(text string) (string, error) {
	f := func(r rune) bool {
		return r == '-' || r == '.'
	}
	if !strings.ContainsFunc(text, f) {
		return morse.ToText(text), nil
	}
	if strings.ContainsFunc(text, f) {
		return morse.ToMorse(text), nil
	}
	return "converting text error", morse.ErrNoEncoding{}
}
