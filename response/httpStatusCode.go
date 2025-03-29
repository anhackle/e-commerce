package response

const (
	ErrCodeSuccess             = 20000
	ErrCodeExternal            = 40000
	ErrCodeNotAuthorize        = 40001
	ErrCodeLoginFail           = 40002
	ErrCodeUserHasExists       = 40003
	ErrCodePasswordNotMatch    = 40004
	ErrCodeOldPasswordNotMatch = 40005
	ErrCodeProductNotFound     = 40006
	ErrCodeQuantityNotEnough   = 40007
	ErrCodeItemNotFoundInCart  = 40008
	ErrCodeCartEmpty           = 40009
	ErrCodeInternal            = 50000
)

// message
var msg = map[int]string{
	ErrCodeSuccess:             "Success",
	ErrCodeLoginFail:           "Username or Password invalid",
	ErrCodeInternal:            "Internal server error",
	ErrCodeExternal:            "Bad request",
	ErrCodeUserHasExists:       "User existed",
	ErrCodeNotAuthorize:        "Not authorized",
	ErrCodePasswordNotMatch:    "Password not match",
	ErrCodeOldPasswordNotMatch: "Old password not match",
	ErrCodeProductNotFound:     "Product not found",
	ErrCodeQuantityNotEnough:   "Stock quantity not enough",
	ErrCodeItemNotFoundInCart:  "Item not found in cart",
	ErrCodeCartEmpty:           "Cart is empty",
}
