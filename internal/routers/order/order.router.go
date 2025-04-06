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
		orderRouterPublic.GET("/", orderController.GetOrders)
		orderRouterPublic.POST("/search", orderController.GetOrder)
		orderRouterPublic.POST("/", orderController.CreateOrder)
		orderRouterPublic.PUT("/pay", orderController.CreatePayment)
		orderRouterPublic.POST("/status", orderController.GetOrderStatus)
	}

	orderRouterPrivate := router.Group("/admin/orders")
	orderRouterPrivate.Use(middlewares.JWTMiddleware())
	orderRouterPrivate.Use(middlewares.RoleVerifyMiddleware())

	{
		orderRouterPrivate.PUT("/status", orderController.UpdateStatus)
		orderRouterPrivate.POST("/", orderController.GetOrdersForAdmin)
		orderRouterPrivate.POST("/search", orderController.GetOrderForAdmin)
		orderRouterPrivate.GET("/summary", orderController.GetOrderSummary)
	}

}
