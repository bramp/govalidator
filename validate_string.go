package govalidator

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrEmpty = errors.New("input is empty")
	ErrNotAInteger = errors.New("input can not be converted to a integer")
)

type stringValidator struct {
	trim  bool
	empty bool
	regex string
}

func NewStringValidator() StringValidator {
	return &stringValidator{false, false, ""}
}

func (v *stringValidator) TrimSpace() StringValidator {
	v.trim = true
	return v
}

func (v *stringValidator) NotEmpty() StringValidator {
	v.empty = true
	return v
}

func (v *stringValidator) Regex(regex string) StringValidator {
	v.regex = regex
	return v
}

func (v *stringValidator) Validate(input string) (string, []error) {
	if v.trim {
		input = strings.TrimSpace(input)
	}

	if v.empty && input == "" {
		return input, []error{ErrEmpty}
	}

	if v.regex != "" {
		panic("regex not supported yet")
	}

	return input, nil
}

/**
 * Attempts to parse the string as a integer
 */
func (v *stringValidator) AsInt() StringIntValidator {
	return &stringIntValidator{
		NewIntValidator(), v,
	}
}

type stringIntValidator struct {
	IntValidator
	StringValidator
}

func (v *stringIntValidator) Range(min, max int) StringIntValidator {
	v.IntValidator.Range(min, max)
	return v
}

func (v *stringIntValidator) Validate(input string) (int, []error) {
	input, errs := v.StringValidator.Validate(input)
	if len(errs) > 0 {
		return 0, errs
	}

	number, err := strconv.Atoi(input)
	if err != nil {
		errs = append(errs, ErrNotAInteger)
		return 0, errs
	}

	return v.IntValidator.Validate(number)
}
