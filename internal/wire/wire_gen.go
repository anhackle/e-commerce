// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"database/sql"
	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/internal/service"
	"github.com/dgraph-io/ristretto/v2"
	"github.com/redis/go-redis/v9"
)

// Injectors from authen.wire.go:

func InitAuthenRouterHandler(dbc *sql.DB) (*controller.AuthenController, error) {
	iAuthenRepo := repo.NewAuthenRepo(dbc)
	iAuthenService := service.NewAuthenService(dbc, iAuthenRepo)
	authenController := controller.NewAuthenController(iAuthenService)
	return authenController, nil
}

// Injectors from cart.wire.go:

func InitCartRouterHandler(dbc *sql.DB) (*controller.CartController, error) {
	iCartRepo := repo.NewCartRepo(dbc)
	iProductRepo := repo.NewProductRepo(dbc)
	iCartService := service.NewCartService(iCartRepo, iProductRepo)
	cartController := controller.NewCartController(iCartService)
	return cartController, nil
}

// Injectors from order.wire.go:

func InitOrderRouterHandler(dbc *sql.DB) (*controller.OrderController, error) {
	iOrderRepo := repo.NewOrderRepo(dbc)
	iCartRepo := repo.NewCartRepo(dbc)
	iProductRepo := repo.NewProductRepo(dbc)
	iOrderService := service.NewOrderService(dbc, iOrderRepo, iCartRepo, iProductRepo)
	orderController := controller.NewOrderController(iOrderService)
	return orderController, nil
}

// Injectors from product.wire.go:

func InitProductRouterHandler(dbc *sql.DB, redisClient *redis.Client, localCache *ristretto.Cache[string, string]) (*controller.ProductController, error) {
	iProductRepo := repo.NewProductRepo(dbc)
	iRedisCache := repo.NewRedisCache(redisClient)
	iLocalCache := repo.NewLocalCache(localCache)
	iProductService := service.NewProductService(iProductRepo, iRedisCache, iLocalCache)
	productController := controller.NewProductController(iProductService)
	return productController, nil
}

// Injectors from user.wire.go:

func InitUserRouterHandler(dbc *sql.DB) (*controller.UserController, error) {
	iUserRepo := repo.NewUserRepo(dbc)
	iUserService := service.NewUserService(iUserRepo)
	userController := controller.NewUserController(iUserService)
	return userController, nil
}
