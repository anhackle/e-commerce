package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidatePaymentMethod(fl validator.FieldLevel) bool {
	var passwordPattern = "^(MOMO|BANK|COD)$"

	password := fl.Field().String()
	match, _ := regexp.MatchString(passwordPattern, password)

	return match
}
