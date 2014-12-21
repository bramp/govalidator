package govalidator

import (
	"testing"
)

func TestIntValidator(t *testing.T) {
	validator := NewIntValidator().Range(0, 10)

	i, errs := validator.Validate(0)
	if i != 0 || len(errs) > 0 {
		t.Errorf("Failed 0 test %v %v", i, errs)
	}

	i, errs = validator.Validate(10)
	if i != 10 || len(errs) > 0 {
		t.Errorf("Failed 10 test %v %v", i, errs)
	}

	i, errs = validator.Validate(-1)
	if len(errs) != 1 || errs[0] != ErrTooSmall {
		t.Errorf("Failed too small test %v %v", i, errs)
	}

	i, errs = validator.Validate(11)
	if len(errs) != 1 || errs[0] != ErrTooLarge {
		t.Errorf("Failed too large test %v %v", i, errs)
	}
}

func TestStringIntValidator(t *testing.T) {
	validator := NewStringValidator().TrimSpace().NotEmpty().AsInt().Range(0, 10)

	i, errs := validator.Validate(" 0")
	if i != 0 || len(errs) > 0 {
		t.Errorf("Failed 0 test %v %v", i, errs)
	}

	i, errs = validator.Validate("10 ")
	if i != 10 || len(errs) > 0 {
		t.Errorf("Failed 10 test %v %v", i, errs)
	}

	i, errs = validator.Validate("-1 ")
	if len(errs) != 1 || errs[0] != ErrTooSmall {
		t.Errorf("Failed too small test %v %v", i, errs)
	}

	i, errs = validator.Validate(" 11")
	if len(errs) != 1 || errs[0] != ErrTooLarge {
		t.Errorf("Failed too large test %v %v", i, errs)
	}

	i, errs = validator.Validate(" ")
	if len(errs) != 1 || errs[0] != ErrEmpty {
		t.Errorf("Failed empty test %v %v", i, errs)
	}
}

func TestMapValidator(t *testing.T) {
	validator := NewMapValidator().
		FailOnUnknown().
		Key("key").NotEmpty().
		Key("int").AsInt().Range(1, 10)

	m, errs := validator.Validate(map[string]interface{}{"key": "hello", "int": "10"})
	if len(errs) != 0 || len(m) != 2 || m["key"] != "hello" || m["int"] != 10 {
		t.Errorf("Failed map test %v %v", m, errs)
	}

	m, errs = validator.Validate(map[string]interface{}{"unknown": "hello"})
	if len(errs) != 1 || errs[0] != ErrUnknownKey {
		t.Errorf("Failed unknown test %v %v", m, errs)
	}
}

func TestReadme(t *testing.T) {

	// Create a map validator, that expects only two fields
	validator := NewMapValidator().
		FailOnUnknown().
		Key("field").NotEmpty().           // TODO FIX THIS
		Key("number").AsInt().Range(1, 10)

	m := map[string]interface{}{
		"field":  "value",
		"number": "10",
	}

	m, errs := validator.Validate(m)
	// Returns {
	//   "field" : "value",
	//   "number": 10,      // Note: int not string.
	// }
	// and errs == nil
	if len(errs) != 0 || len(m) != 2 || m["field"] != "value" || m["number"] != 10 {
		t.Errorf("Failed map test %v %v", m, errs)
	}

	// However
	m = map[string]interface{}{
		"field":  "",
		"number": "1000",
		"other":  "value",
	}
	m, errs = validator.Validate(m)
	// Returns {
	//   "field" : "",
	//   "number": 1000,
	// }
	// and errs == {ErrEmpty, ErrTooLarge, ErrUnknownKey}
	if len(errs) != 3 || len(m) != 2 || m["field"] != "" || m["number"] != 1000 {
		t.Errorf("Failed map test %v %v", m, errs)
	}
}
