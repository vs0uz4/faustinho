package utils

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyURL         = errors.New("URL cannot be empty")
	ErrInvalidURLFormat = errors.New("Invalid URL format")
)

func ErrNumberMustBePositive(name string) error {
	return fmt.Errorf("%s must be greater than 0", name)
}
