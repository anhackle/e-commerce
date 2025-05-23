package initialize

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/middlewares"
	"github.com/anle/codebase/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	authenRouter := routers.RouterGroupApp.Authen
	userRouter := routers.RouterGroupApp.User
	productRouter := routers.RouterGroupApp.Product
	cartRouter := routers.RouterGroupApp.Cart
	orderRouter := routers.RouterGroupApp.Order

	MainGroup := r.Group("/v1")
	MainGroup.Use(middlewares.CORSMiddleware())

	{
		authenRouter.InitAuthenRouter(MainGroup)
		userRouter.InitUserRouter(MainGroup)
		productRouter.InitProductRouter(MainGroup)
		cartRouter.InitCartRouter(MainGroup)
		orderRouter.InitOrderRouter(MainGroup)
	}

	return r
}
