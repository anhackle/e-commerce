package routers

import (
	"github.com/anle/codebase/internal/routers/authen"
	"github.com/anle/codebase/internal/routers/cart"
	"github.com/anle/codebase/internal/routers/order"
	"github.com/anle/codebase/internal/routers/product"
	"github.com/anle/codebase/internal/routers/user"
)

type RouterGroup struct {
	Authen  authen.AuthenRouterGroup
	User    user.UserRouterGroup
	Product product.ProductRouterGroup
	Cart    cart.CartRouterGroup
	Order   order.OrderRouterGroup
}

var RouterGroupApp = new(RouterGroup)
