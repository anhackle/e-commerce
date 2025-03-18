package response

const (
	ErrCodeSuccess             = 20000
	ErrCodeExternal            = 40000
	ErrCodeNotAuthorize        = 40001
	ErrCodeLoginFail           = 40002
	ErrCodeUserHasExists       = 40003
	ErrCodePasswordNotMatch    = 40004
	ErrCodeOldPasswordNotMatch = 40005
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
}
