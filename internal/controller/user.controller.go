package controller

import (
	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/service"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func (uc *UserController) GetProfile(c *gin.Context) {
	profileResult, result, _ := uc.userService.GetProfile(c)

	response.HandleResult(c, result, profileResult)
}

func (uc *UserController) UpdateProfile(c *gin.Context) {
	var input model.UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := uc.userService.UpdateProfile(c, input)

	response.HandleResult(c, result, nil)
}

func (uc *UserController) ChangePassword(c *gin.Context) {
	var input model.ChangePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := uc.userService.ChangePassword(c, input)

	response.HandleResult(c, result, nil)
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
