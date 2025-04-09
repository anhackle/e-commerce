package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateUUID(fl validator.FieldLevel) bool {
	var uuidPattern = "^[a-z0-9\\-]+$"

	uuid := fl.Field().String()

	match, _ := regexp.MatchString(uuidPattern, uuid)

	return match
}
