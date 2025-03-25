//go:build wireinject

package wire

import (
	"database/sql"

	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/internal/service"
	"github.com/google/wire"
)

func InitCartRouterHandler(dbc *sql.DB) (*controller.CartController, error) {
	wire.Build(
		repo.NewCartRepo,
		repo.NewProductRepo,
		service.NewCartService,
		controller.NewCartController,
	)

	return new(controller.CartController), nil
}
