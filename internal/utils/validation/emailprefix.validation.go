package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateEmailPrefix(fl validator.FieldLevel) bool {
	var emailPrefixPattern = "^[a-zA-Z0-9._+\\-@]*$"

	emailPrefix := fl.Field().String()
	if emailPrefix == "" {
		return true
	}

	match, _ := regexp.MatchString(emailPrefixPattern, emailPrefix)

	return match
}
