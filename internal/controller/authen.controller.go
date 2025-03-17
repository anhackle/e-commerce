package controller

import (
	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/service"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type AuthenController struct {
	authenService service.IAuthenService
}

func (ac *AuthenController) Register(c *gin.Context) {
	var input model.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := ac.authenService.Register(c, input)
	response.HandleResult(c, result, nil)
}

func (ac *AuthenController) Login(c *gin.Context) {
	var input model.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	token, result, _ := ac.authenService.Login(c, input)
	response.HandleResult(c, result, token)
}

func NewAuthenController(authenService service.IAuthenService) *AuthenController {
	return &AuthenController{
		authenService: authenService,
	}
}
