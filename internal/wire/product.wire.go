//go:build wireinject

package wire

import (
	"database/sql"

	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/internal/service"
	"github.com/dgraph-io/ristretto/v2"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

func InitProductRouterHandler(dbc *sql.DB, redisClient *redis.Client, localCache *ristretto.Cache[string, string]) (*controller.ProductController, error) {
	wire.Build(
		repo.NewLocalCache,
		repo.NewRedisCache,
		repo.NewProductRepo,
		service.NewProductService,
		controller.NewProductController,
	)

	return new(controller.ProductController), nil
}
