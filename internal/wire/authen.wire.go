//go:build wireinject

package wire

import (
	"database/sql"

	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/internal/service"
	"github.com/google/wire"
)

func InitAuthenRouterHandler(dbc *sql.DB) (*controller.AuthenController, error) {
	wire.Build(
		repo.NewAuthenRepo,
		service.NewAuthenService,
		controller.NewAuthenController,
	)

	return new(controller.AuthenController), nil
}
