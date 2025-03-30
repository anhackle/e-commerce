package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateStatus(fl validator.FieldLevel) bool {
	var statusPattern = `^(create|confirm|pay|ship|finish)$`

	status := fl.Field().String()
	match, _ := regexp.MatchString(statusPattern, status)

	return match
}
