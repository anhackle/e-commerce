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
		UserRouterPublic.GET("/profile", userController.GetProfile)
		UserRouterPublic.PUT("/profile", userController.UpdateProfile)
		UserRouterPublic.PUT("/password", userController.ChangePassword)
	}

	UserRouterPrivate := router.Group("/admin/user")
	UserRouterPrivate.Use(middlewares.JWTMiddleware())
	UserRouterPrivate.Use(middlewares.RoleVerifyMiddleware())
	{
		UserRouterPrivate.PUT("/role", userController.UpdateRole)
	}
}
