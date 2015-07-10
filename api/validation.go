package api

import validator "github.com/asaskevich/govalidator"

func init() {
	validator.TagMap["grant_type"] = validator.Validator(isGrantType)
}

func isGrantType(s string) bool {
	return s == "authorization_code" || s == "implicit" || s == "password" || s == "refresh_token"
}
