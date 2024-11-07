package user

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/service"
	"github.com/anle/codebase/model"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthenController struct{}

var Authen = new(AuthenController)

func (c *AuthenController) VerifyOTP(ctx *gin.Context) {
	var params model.VerifyInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponseExternal(ctx, response.ErrCodeExternal, err.Error())
		return
	}

	result, err := service.UserAuthen().VerifyOTP(ctx, &params)
	if err != nil {
		response.ErrorResponseExternal(ctx, response.ErrTokenInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}
func (c *AuthenController) Login(ctx *gin.Context) {
	err := service.UserAuthen().Login(ctx)
	if err != nil {
		response.ErrorResponseExternal(ctx, response.ErrCodeExternal, nil)
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}

// ShowAccount godoc
// @Summary      User Registration
// @Description  When user registered, send otp to email
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        payload body model.RegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       //users/authen/register [post]
func (c *AuthenController) Register(ctx *gin.Context) {
	var params model.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponseExternal(ctx, response.ErrCodeExternal, err.Error())
		return
	}

	codeStatus, err := service.UserAuthen().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Error registering user OTP", zap.Error(err))
		response.ErrorResponseInternal(ctx, response.ErrCodeInternal, err)
		return
	}

	response.SuccessResponse(ctx, codeStatus, nil)
}
