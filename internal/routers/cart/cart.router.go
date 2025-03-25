package cart

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/middlewares"
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type CartRouter struct{}

func (cr *CartRouter) InitCartRouter(router *gin.RouterGroup) {
	cartController, _ := wire.InitCartRouterHandler(global.Mdb)

	cartRouterPublic := router.Group("/cart")
	cartRouterPublic.Use(middlewares.JWTMiddleware())

	{
		cartRouterPublic.GET("/", cartController.GetCart)
		cartRouterPublic.POST("/", cartController.AddToCart)
		cartRouterPublic.DELETE("/", cartController.DeleteCart)
	}

}
