//go:build wireinject

package wire

import (
	"database/sql"

	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/internal/service"
	"github.com/google/wire"
)

func InitOrderRouterHandler(dbc *sql.DB) (*controller.OrderController, error) {
	wire.Build(
		repo.NewOrderRepo,
		repo.NewCartRepo,
		repo.NewProductRepo,
		service.NewOrderService,
		controller.NewOrderController,
	)

	return new(controller.OrderController), nil
}
