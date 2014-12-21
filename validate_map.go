package govalidator

import (
	"errors"
)

var (
	ErrUnknownKey = errors.New("unknown key")
)

type mapValidator struct {
	validators    map[string]interfaceValidator
	failOnUnknown bool
}

type mapStringValidator struct {
	root      *mapValidator
	key       string
	validator StringValidator
}

type mapStringIntValidator struct {
	root      *mapValidator
	key       string
	validator StringIntValidator
}

func NewMapValidator() MapValidator {
	return &mapValidator{
		map[string]interfaceValidator{},
		false,
	}
}

func (v *mapValidator) FailOnUnknown() MapValidator {
	v.failOnUnknown = true
	return v
}

func (v *mapValidator) Key(key string) MapStringValidator {
	validator := &mapStringValidator{v, key, NewStringValidator()}
	v.validators[key] = validator
	return validator
}

// For strings
func (v *mapStringValidator) TrimSpace() MapStringValidator {
	v.validator = v.validator.TrimSpace()
	return v
}

func (v *mapStringValidator) NotEmpty() MapStringValidator {
	v.validator = v.validator.NotEmpty()
	return v
}

func (v *mapStringValidator) Regex(regex string) MapStringValidator {
	v.validator = v.validator.Regex(regex)
	return v
}

func (v *mapStringValidator) AsInt() MapStringIntValidator {
	validator := &mapStringIntValidator{v.root, v.key, v.validator.AsInt()}
	validator.root.validators[v.key] = validator
	return validator
}

func (v *mapStringValidator) validateInterface(input interface{}) (interface{}, []error) {
	return v.validator.Validate(input.(string))
}

func (v *mapStringValidator) Key(key string) MapStringValidator {
	return v.root.Key(key)
}

func (v *mapStringValidator) Validate(input map[string]interface{}) (map[string]interface{}, []error) {
	return v.root.Validate(input)
}

func (v *mapStringIntValidator) Range(min, max int) MapStringIntValidator {
	v.validator = v.validator.Range(min, max)
	return v
}

func (v *mapStringIntValidator) validateInterface(input interface{}) (interface{}, []error) {
	return v.validator.Validate(input.(string))
}

func (v *mapStringIntValidator) Key(key string) MapStringValidator {
	return v.root.Key(key)
}

func (v *mapStringIntValidator) Validate(input map[string]interface{}) (map[string]interface{}, []error) {
	return v.root.Validate(input)
}

func (v *mapValidator) Validate(input map[string]interface{}) (map[string]interface{}, []error) {
	errs := make([]error, 0)
	for key, value := range input {
		validator, found := v.validators[key]
		if found {
			value, new_errs := validator.validateInterface(value)

			input[key] = value
			errs = append(errs, new_errs...)

		} else if v.failOnUnknown {
			errs = append(errs, ErrUnknownKey)
			delete(input, key)
		}
	}

	return input, errs
}


