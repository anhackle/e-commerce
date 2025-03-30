package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidatePaymentMethod(fl validator.FieldLevel) bool {
	var paymentPattern = "^(MOMO|BANK|COD)$"

	payment := fl.Field().String()
	match, _ := regexp.MatchString(paymentPattern, payment)

	return match
}
