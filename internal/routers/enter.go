package routers

import (
	"github.com/anle/codebase/internal/routers/admin"
	"github.com/anle/codebase/internal/routers/authen"
	"github.com/anle/codebase/internal/routers/user"
)

type RouterGroup struct {
	Authen authen.AuthenRouterGroup
	User   user.UserRouterGroup
	Admin  admin.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)
