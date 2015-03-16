package validator

import (
	"bytes"
	"errors"
	"reflect"
	"regexp"
	"xavier/lib/oauth"

	"github.com/mccoyst/validate"
)

var (
	errValidationEmail            = errors.New("Should be a valid email address")
	errValidationGrantType        = errors.New("Should be a valid grant type")
	errValidationNonZero          = errors.New("Should not be empty or null")
	errValidationPasswordTooShort = errors.New("Should contain 6 or more characters")
	errValidationPasswordTooLong  = errors.New("Should contain 128 or less characters")
	errValidationUUID             = errors.New("Not a valid UUID")
)

var emailRegexp = regexp.MustCompile("^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")
var uuidRegexp = regexp.MustCompile("^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$")

type Validator struct {
	validator *validate.V
}

type ValidationError struct {
	Errors []error `json:"errors,omitempty"`
}

func New() *Validator {
	v := make(validate.V)
	v["email"] = email
	v["grant_type"] = grantType
	v["nonzero"] = nonZero
	v["password"] = password
	v["uuid"] = uuid
	return &Validator{&v}
}

func (v *Validator) Validate(obj interface{}) *ValidationError {
	e := v.validator.Validate(obj)
	if len(e) > 0 {
		return &ValidationError{e}
	}

	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	if typ.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			e := v.validator.Validate(val.Index(i).Interface())
			if len(e) > 0 {
				return &ValidationError{e}
			}
		}
	}
	return nil
}

func (e *ValidationError) Error() string {
	var buffer bytes.Buffer
	for _, err := range e.Errors {
		buffer.WriteString(err.Error())
	}
	return buffer.String()
}

// Private

func email(i interface{}) error {
	s := i.(string)
	if !emailRegexp.MatchString(s) {
		return errValidationEmail
	}
	return nil
}

func grantType(i interface{}) error {
	s := i.(string)
	if !oauth.ValidGrantType(s) {
		return errValidationGrantType
	}
	return nil
}

func nonZero(i interface{}) error {
	s := i.(string)
	if len(s) == 0 {
		return errValidationNonZero
	}
	return nil
}

func password(i interface{}) error {
	s := i.(string)
	length := len([]rune(s))
	if length < 6 {
		return errValidationPasswordTooShort
	} else if length > 128 {
		return errValidationPasswordTooLong
	}
	return nil
}

func uuid(i interface{}) error {
	s := i.(string)
	if !uuidRegexp.MatchString(s) {
		return errValidationUUID
	}
	return nil
}
