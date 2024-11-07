package initialize

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/service"
	"github.com/anle/codebase/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	// User Service Interface
	service.InitUserAuthen(impl.NewUserAuthenImpl(queries))
	// .......

}
