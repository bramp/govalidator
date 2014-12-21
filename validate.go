package govalidator

type IntValidator interface {
	Range(min, max int) IntValidator

	Validate(input int) (int, []error)
}

// NB: I don't like this interface, and open to suggestions on a way of removing it.
type StringIntValidator interface {
	Range(min, max int) StringIntValidator

	Validate(input string) (int, []error)
}

type StringValidator interface {
	TrimSpace() StringValidator
	NotEmpty() StringValidator
	Regex(regex string) StringValidator
	AsInt() StringIntValidator

	Validate(input string) (string, []error)
}

type interfaceValidator interface {
	validateInterface(input interface{}) (interface{}, []error)
}

type MapValidator interface {
	FailOnUnknown() MapValidator

	Key(key string) MapStringValidator
	// TODO Add IntKey(key string) MapIntValidator

	Validate(input map[string]interface{}) (map[string]interface{}, []error)
}

type MapStringValidator interface {
	Key(key string) MapStringValidator
	// TODO Add IntKey(key string) MapIntValidator

	TrimSpace() MapStringValidator
	NotEmpty() MapStringValidator
	Regex(regex string) MapStringValidator
	AsInt() MapStringIntValidator

	Validate(input map[string]interface{}) (map[string]interface{}, []error)
}

type MapStringIntValidator interface {
	Key(key string) MapStringValidator
	// TODO Add IntKey(key string) MapIntValidator

	Range(min, max int) MapStringIntValidator

	Validate(input map[string]interface{}) (map[string]interface{}, []error)
}
