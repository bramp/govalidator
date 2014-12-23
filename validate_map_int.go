package govalidator

type mapIntValidator struct {
	mapValidatorEntry
	validator IntValidator
}

func NewMapIntValidator(root *mapValidator, key string) *mapIntValidator {
	return &mapIntValidator{
		NewMapValidatorEntry(root, key),
		NewIntValidator(),
	}
}

func (v *mapIntValidator) Range(min, max int) MapIntValidator {
	v.validator = v.validator.Range(min, max)
	return v
}

func (v *mapIntValidator) validate(input interface{}) (interface{}, []error) {
	s, ok := input.(int)
	if ok {
		return v.validator.Validate(s)
	}
	return input, []error{ErrNotAInteger}
}
