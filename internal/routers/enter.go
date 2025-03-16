package routers

import "github.com/anle/codebase/internal/routers/authen"

type RouterGroup struct {
	Authen authen.AuthenRouter
}

var RouterGroupApp = new(RouterGroup)
