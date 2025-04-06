package product

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/middlewares"
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(router *gin.RouterGroup) {
	productController, _ := wire.InitProductRouterHandler(global.Mdb)

	productRouterPublic := router.Group("/products")

	{
		productRouterPublic.POST("/search", productController.GetProducts)
	}

	productRouterPrivate := router.Group("/admin/products")
	productRouterPrivate.Use(middlewares.JWTMiddleware())
	productRouterPrivate.Use(middlewares.RoleVerifyMiddleware())

	{
		productRouterPrivate.POST("/search", productController.GetProductsForAdmin)
		productRouterPrivate.POST("/", productController.CreateProduct)
		productRouterPrivate.PUT("/", productController.UpdateProduct)
		productRouterPrivate.DELETE("/", productController.DeleteProduct)
	}
}
