package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateOrderBy(fl validator.FieldLevel) bool {
	var orderByPattern = "^(created_at|total)$"

	orderBy := fl.Field().String()
	if orderBy == "" {
		return true
	}

	match, _ := regexp.MatchString(orderByPattern, orderBy)

	return match
}
