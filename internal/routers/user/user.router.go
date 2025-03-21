package user

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/middlewares"
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler(global.Mdb)

	UserRouterPublic := router.Group("/user")
	UserRouterPublic.Use(middlewares.JWTMiddleware())

	{
		UserRouterPublic.PUT("/profile", userController.UpdateProfile)
		UserRouterPublic.POST("/password", userController.ChangePassword)
	}
}
