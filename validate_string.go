package govalidator

import (
	"strings"
)

type stringValidator struct {
	trim  bool
	empty bool
	regex string
	fun   StringValidatorFunc
}

func NewStringValidator() StringValidator {
	return &stringValidator{false, false, "", nil}
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

func (v *stringValidator) Func(fun StringValidatorFunc) StringValidator {
	v.fun = fun
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

	if v.fun != nil {
		return v.fun(input)
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

/**
 * Attempts to parse the string as a bool
 */
func (v *stringValidator) AsBool() StringBoolValidator {
	return &stringBoolValidator{
		v, "true", "false", nil,
	}
}
