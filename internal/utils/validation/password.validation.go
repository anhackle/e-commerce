package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidatePassword(fl validator.FieldLevel) bool {
	var passwordPattern = "^[a-zA-Z0-9!@#$%^&*()]+$"

	password := fl.Field().String()
	match, _ := regexp.MatchString(passwordPattern, password)

	return match
}
