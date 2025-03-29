package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidatePhone(fl validator.FieldLevel) bool {
	var phonePattern = `^(03[2-9]|05[6|8|9]|07[0-9]|08[1-6|8|9]|09[0-9])\d{7}$`

	phone := fl.Field().String()
	match, _ := regexp.MatchString(phonePattern, phone)

	return match
}
