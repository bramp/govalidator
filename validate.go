package govalidator

type IntValidatorFunc func (input int) (output int, errors []error)
type IntValidate interface {
	Validate(input int) (output int, errors []error)
}

type IntValidator interface {
	Range(min, max int) IntValidator

	IntValidate
}

type StringValidatorFunc func (input string) (output string, errors []error)
type StringValidatorCommon interface {
	Validate(input string) (output string, errors []error)
}

type StringValidator interface {

	TrimSpace() StringValidator
	NotEmpty() StringValidator
	Regex(regex string) StringValidator
	Func(f StringValidatorFunc) StringValidator

	AsInt() StringIntValidator
	AsBool() StringBoolValidator

	StringValidatorCommon
}

type StringIntValidator interface {
	Range(min, max int) StringIntValidator

	Validate(input string) (int, []error)
}

type StringBoolValidator interface {
	True(true string) StringBoolValidator
	False(false string) StringBoolValidator

	Validate(input string) (bool, []error)
}

type genericMapValidator interface {
	validate(input interface{}) (interface{}, []error)
	validateMissing() (interface{}, []error)
}

type MapValidatorCommon interface {

	Key(key string) MapStringValidator
	IntKey(key string) MapIntValidator

	Validate(input map[string]interface{}) (map[string]interface{}, map[string][]error)
}

type MapValidator interface {
	MapValidatorCommon
	FailOnUnknown() MapValidator
}

type MapIntValidator interface {
	MapValidatorCommon

	Range(min, max int) MapIntValidator
}

type MapStringValidator interface {
	MapValidatorCommon

	Required() MapStringValidator
	Default(defaut interface{}) MapStringValidator

	AsInt() MapStringIntValidator
	AsBool() MapStringBoolValidator

	TrimSpace() MapStringValidator
	NotEmpty() MapStringValidator
	Regex(regex string) MapStringValidator
	Func(f StringValidatorFunc) MapStringValidator
}

type MapStringIntValidator interface {
	MapValidatorCommon

	Range(min, max int) MapStringIntValidator
}

type MapStringBoolValidator interface {
	MapValidatorCommon

	True(true string) MapStringBoolValidator
	False(false string) MapStringBoolValidator
}
