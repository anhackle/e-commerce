package user

import (
	"github.com/anle/codebase/internal/controller/user"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (p *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	authenRouterPublic := router.Group("/users/authen")

	{
		authenRouterPublic.POST("/login", user.Authen.Login)
		authenRouterPublic.POST("/register", user.Authen.Register)
	}

}
