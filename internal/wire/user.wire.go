//go:build wireinject

package wire

import (
	"database/sql"

	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHandler(dbc *sql.DB) (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepo,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
