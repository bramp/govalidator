package govalidator

type mapStringIntValidator struct {
	mapValidatorEntry
	validator StringIntValidator
}

func NewMapStringIntValidator(entry mapValidatorEntry, validator StringIntValidator) *mapStringIntValidator {
	return &mapStringIntValidator{
		entry,
		validator,
	}
}

func (v *mapStringIntValidator) Range(min, max int) MapStringIntValidator {
	v.validator = v.validator.Range(min, max)
	return v
}

func (v *mapStringIntValidator) Key(key string) MapStringValidator {
	return v.root.Key(key)
}

func (v *mapStringIntValidator) validate(input interface{}) (interface{}, []error) {
	s, ok := input.(string)
	if ok {
		return v.validator.Validate(s)
	}
	return input, []error{ErrNotAString}
}
