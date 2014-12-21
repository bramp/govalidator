package govalidator

import (
	"errors"
)

var (
	ErrTooSmall = errors.New("input is too small")
	ErrTooLarge = errors.New("input is too large")
)

// Note the use of << to create an untyped constant.
const bitsPerWord = 32 << uint(^uint(0)>>63)

// Implementation-specific integer limit values.
// Taken from: http://code.google.com/p/go-bit/
const (
	maxInt  = 1<<(bitsPerWord-1) - 1 // either 1<<31 - 1 or 1<<63 - 1
	minInt  = -maxInt - 1            // either -1 << 31 or -1 << 63
	maxUint = 1<<bitsPerWord - 1     // either 1<<32 - 1 or 1<<64 - 1
	minUint = 0
)

type intValidator struct {
	min, max int
}

func NewIntValidator() IntValidator {
	return &intValidator{
		min: minInt,
		max: maxInt,
	}
}

func (v *intValidator) Range(min, max int) IntValidator {
	v.min = min
	v.max = max

	return v
}

func (v *intValidator) Validate(input int) (int, []error) {
	if input < v.min {
		return input, []error{ErrTooSmall}
	}
	if input > v.max {
		return input, []error{ErrTooLarge}
	}
	return input, nil
}
