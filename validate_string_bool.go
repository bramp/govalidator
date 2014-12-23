package govalidator

type stringBoolValidator struct {
	StringValidator
	true   string
	false  string
	defaut *bool
}

func (v *stringBoolValidator) True(t string) StringBoolValidator {
	v.true = t
	return v
}

func (v *stringBoolValidator) False(f string) StringBoolValidator {
	v.false = f
	return v
}

func (v *stringBoolValidator) Default(defaut bool) StringBoolValidator {
	v.defaut = &defaut
	return v
}

func (v *stringBoolValidator) Validate(input string) (bool, []error) {
	if input == v.true {
		return true, nil
	}
	if input == v.false {
		return false, nil
	}
	if v.defaut != nil {
		return *v.defaut, nil
	}

	return false, []error{ErrNotABoolean}
}
