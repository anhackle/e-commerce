//go:build wireinject

package wire

import (
	"database/sql"

	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/internal/service"
	"github.com/google/wire"
)

func InitProductRouterHandler(dbc *sql.DB) (*controller.ProductController, error) {
	wire.Build(
		repo.NewProductRepo,
		service.NewProductService,
		controller.NewProductController,
	)

	return new(controller.ProductController), nil
}
