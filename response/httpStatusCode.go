package response

const (
	ErrCodeSuccess              = 20000
	ErrCodeExternal             = 40000
	ErrCodeParamInvalid         = 40001
	ErrCodeLoginFail            = 40002
	ErrCodeUserHasExists        = 40003
	ErrCodeNegativeAmount       = 40004
	ErrCodeNotEnoughBalance     = 40005
	ErrCodeAccountNotExist      = 40006
	ErrCodeGroupNotExist        = 40007
	ErrCodePercentageExceed     = 40008
	ErrCodeDistributionNotExist = 40009
	ErrCodeGroupNotFull         = 40010
	ErrCodeOTPInvalid           = 40011
	ErrTokenInvalid             = 40100
	ErrCodeInternal             = 50000
)

// message
var msg = map[int]string{
	ErrCodeSuccess:              "Success",
	ErrCodeLoginFail:            "Username or Password invalid",
	ErrCodeInternal:             "Internal server error",
	ErrCodeExternal:             "Bad request",
	ErrCodeParamInvalid:         "Email is invalid",
	ErrCodeNegativeAmount:       "Amount must be greater than 0",
	ErrCodeNotEnoughBalance:     "Balance not enough",
	ErrCodeAccountNotExist:      "Account not exist",
	ErrCodeGroupNotExist:        "Group not exist",
	ErrCodePercentageExceed:     "Percentage Exceed",
	ErrCodeDistributionNotExist: "Account not in group",
	ErrCodeGroupNotFull:         "Group not full yet",
	ErrCodeOTPInvalid:           "OTP invalid",
	ErrTokenInvalid:             "Authorization required",
	ErrCodeUserHasExists:        "User existed",
}
