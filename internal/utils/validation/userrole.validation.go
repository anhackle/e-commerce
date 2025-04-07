package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateUserRole(fl validator.FieldLevel) bool {
	var rolePattern = "^(admin|user)$"

	role := fl.Field().String()
	if role == "" {
		return true
	}

	match, _ := regexp.MatchString(rolePattern, role)

	return match
}
