package order

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/middlewares"
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct{}

func (cr *OrderRouter) InitOrderRouter(router *gin.RouterGroup) {
	orderController, _ := wire.InitOrderRouterHandler(global.Mdb)

	orderRouterPublic := router.Group("/orders")
	orderRouterPublic.Use(middlewares.JWTMiddleware())

	{
		orderRouterPublic.GET("/", orderController.GetOrder)
		orderRouterPublic.POST("/", orderController.CreateOrder)
		orderRouterPublic.DELETE("/", orderController.DeleteOrder)
	}

}
