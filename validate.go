package govalidator

type IntValidator interface {
	Range(min, max int) IntValidator

	Validate(input int) (int, []error)
}

type StringValidator interface {
	TrimSpace() StringValidator
	NotEmpty() StringValidator
	Regex(regex string) StringValidator

	AsInt() StringIntValidator
	AsBool() StringBoolValidator

	Validate(input string) (string, []error)
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

type MapValidator interface {
	FailOnUnknown() MapValidator

	Key(key string) MapStringValidator
	// TODO Add IntKey(key string) MapIntValidator

	Validate(input map[string]interface{}) (map[string]interface{}, map[string][]error)
}

type MapStringValidator interface {
	Key(key string) MapStringValidator
	// TODO Add IntKey(key string) MapIntValidator

	Required() MapStringValidator
	Default(defaut interface{}) MapStringValidator

	AsInt() MapStringIntValidator
	AsBool() MapStringBoolValidator

	TrimSpace() MapStringValidator
	NotEmpty() MapStringValidator
	Regex(regex string) MapStringValidator

	Validate(input map[string]interface{}) (map[string]interface{}, map[string][]error)
}

type MapStringIntValidator interface {
	Key(key string) MapStringValidator
	// TODO Add IntKey(key string) MapIntValidator

	Range(min, max int) MapStringIntValidator

	Validate(input map[string]interface{}) (map[string]interface{}, map[string][]error)
}

type MapStringBoolValidator interface {
	Key(key string) MapStringValidator
	// TODO Add IntKey(key string) MapIntValidator

	True(true string) MapStringBoolValidator
	False(false string) MapStringBoolValidator

	Validate(input map[string]interface{}) (map[string]interface{}, map[string][]error)
}
