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

func (uc *UserController) UpdateRole(c *gin.Context) {
	var input model.UpdateRoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := uc.userService.UpdateRole(c, input)

	response.HandleResult(c, result, nil)
}

func (uc *UserController) GetUsersForAdmin(c *gin.Context) {
	var input model.GetUsersForAdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	users, result, _ := uc.userService.GetUsersForAdmin(c, input)

	response.HandleResult(c, result, users)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	var input model.DeleteUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := uc.userService.DeleteUser(c, input)

	response.HandleResult(c, result, nil)
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
