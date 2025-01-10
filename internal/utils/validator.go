package utils

import (
	"net/url"
)

func ValidateURL(input string) error {
	if input == "" {
		return ErrEmptyURL
	}
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return ErrInvalidURLFormat
	}
	return nil
}

func ValidatePositiveNumber(value int, name string) error {
	if value <= 0 {
		return ErrNumberMustBePositive(name)
	}
	return nil
}
