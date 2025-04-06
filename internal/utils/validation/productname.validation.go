package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateProductName(fl validator.FieldLevel) bool {
	var namePattern = "^[\\p{L} \\d]+$"

	name := fl.Field().String()
	match, _ := regexp.MatchString(namePattern, name)

	return match
}
