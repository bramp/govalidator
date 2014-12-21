govalidator
============
2014 by [Andrew Brampton](http://bramp.net)

Simple input validation for Google Go / Golang. Use the below examples, or read the [godocs](https://godoc.org/github.com/bramp/govalidator).

Install
-------
```bash
$ go get github.com/bramp/govalidator
```

```go
import (
    "github.com/bramp/govalidator"
)
```

String Example
-------
```go
validator := NewStringValidator().TrimSpace().NotEmpty()
s, errs := validator.Validate(" blah ")
// s == "blah" and errs == nil

s, errs := validator.Validate("  ")
// s == "" and errs = {ErrEmpty}
```

String to Integer Example
-------
```go
validator := NewStringValidator().AsInt().Range(0, 10)
i, errs := validator.Validate("10")
// i == 10, and errs == nil

i, errs := validator.Validate("foo")
// errs = {ErrNotAInteger}

i, errs := validator.Validate("1000")
// errs = {ErrTooLarge}
```

Map Example
-------
```go
// Create a map validator, that expects only two fields
validator := NewMapValidator().
    FailOnUnknown().
    Key("field").NotEmpty().
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
```

