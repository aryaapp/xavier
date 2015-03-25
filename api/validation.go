package api

import (
	"reflect"
	"regexp"
	"xavier/lib/oauth"

	validator "github.com/asaskevich/govalidator"
)

var uuidRegexp = regexp.MustCompile("^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$")

func init() {
	validator.TagMap["grant_type"] = validator.Validator(isGrantType)
}

func Validate(obj interface{}) error {
	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	if typ.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			_, err := validator.ValidateStruct(val.Index(i).Interface())
			if err != nil {
				return err
			}
		}
	} else {
		_, err := validator.ValidateStruct(obj)
		return err
	}
	return nil
}

func isGrantType(s string) bool {
	return oauth.ValidGrantType(s)
}

func isMinimumPassword(s string) bool {
	length := len([]rune(s))
	return length < 6
	// return length < 6 {
	// 	return errValidationPasswordTooShort
	// } else if length > 128 {
	// 	return errValidationPasswordTooLong
	// }
	// return nil
}
