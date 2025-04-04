package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateStatus(fl validator.FieldLevel) bool {
	var statusPattern = `^(pending|paid|processing|shipped|delivered|cancelled|failed)$`

	status := fl.Field().String()
	match, _ := regexp.MatchString(statusPattern, status)

	return match
}
