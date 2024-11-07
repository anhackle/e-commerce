package initialize

import (
	"github.com/anle/codebase/internal/utils/validation"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//Register custom validation here
		v.RegisterValidation("password", validation.ValidatePassword)
	}
}
