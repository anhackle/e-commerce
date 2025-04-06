package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateImageExtension(fl validator.FieldLevel) bool {
	var imagePattern = "^.+\\.(jpg|png|gif)$"

	imageext := fl.Field().String()
	match, _ := regexp.MatchString(imagePattern, imageext)

	return match
}
