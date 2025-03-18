package routers

import (
	"github.com/anle/codebase/internal/routers/authen"
	"github.com/anle/codebase/internal/routers/user"
)

type RouterGroup struct {
	Authen authen.AuthenRouter
	User   user.UserRouter
}

var RouterGroupApp = new(RouterGroup)
