package govalidator

import (
	"strconv"
)

type stringIntValidator struct {
	IntValidator
	StringValidator
}

func (v *stringIntValidator) Range(min, max int64) StringIntValidator {
	v.IntValidator.Range(min, max)
	return v
}

func (v *stringIntValidator) Validate(input string) (int64, []error) {
	input, errs := v.StringValidator.Validate(input)
	if len(errs) > 0 {
		return 0, errs
	}

	number, err := strconv.ParseInt(input, 10, 0)
	if err != nil {
		errs = append(errs, ErrNotAInteger)
		return 0, errs
	}

	return v.IntValidator.Validate(number)
}
