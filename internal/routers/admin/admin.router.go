package admin

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/middlewares"
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (ur *AdminRouter) InitAdminRouter(router *gin.RouterGroup) {
	adminController, _ := wire.InitAdminRouterHandler(global.Mdb)

	adminRouterPublic := router.Group("/admin")
	adminRouterPublic.Use(middlewares.JWTMiddleware())

	{
		adminRouterPublic.GET("/users", adminController.UpdateProfile)
		adminRouterPublic.GET("/users/:id", adminController.ChangePassword)
	}
}
