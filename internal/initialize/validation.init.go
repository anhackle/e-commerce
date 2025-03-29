package initialize

import (
	"github.com/anle/codebase/internal/utils/validation"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", validation.ValidatePassword)
		v.RegisterValidation("name", validation.ValidateName)
		v.RegisterValidation("phone", validation.ValidatePhone)
		v.RegisterValidation("paymentmethod", validation.ValidatePaymentMethod)
	}
}
