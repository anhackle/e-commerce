package routers

import (
	"github.com/anle/codebase/internal/routers/authen"
	"github.com/anle/codebase/internal/routers/product"
	"github.com/anle/codebase/internal/routers/user"
)

type RouterGroup struct {
	Authen  authen.AuthenRouter
	User    user.UserRouter
	Product product.ProductRouter
}

var RouterGroupApp = new(RouterGroup)
