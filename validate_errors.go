package govalidator

import "errors"

var (
	ErrTooSmall = errors.New("input is too small")
	ErrTooLarge = errors.New("input is too large")

	ErrEmpty       = errors.New("input is empty")
	ErrNotAString  = errors.New("input is not a string")
	ErrNotAInteger = errors.New("input is not a integer")
	ErrNotABoolean = errors.New("input is not a boolean")

	ErrRequiredKeyMissing = errors.New("required key missing")
	ErrUnknownKey         = errors.New("unknown key")
)
