package authen

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type AuthenRouter struct{}

func (ar *AuthenRouter) InitAuthenRouter(router *gin.RouterGroup) {
	authenController, _ := wire.InitAuthenRouterHandler(global.Mdb)

	AuthenRouterPublic := router.Group("/authen")

	{
		AuthenRouterPublic.POST("/register", authenController.Register)
	}
}
