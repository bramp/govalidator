package govalidator

import (
	"strconv"
)

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
