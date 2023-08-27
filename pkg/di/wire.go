//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/munsheerck79/Ecom_project.git/pkg/api"
	"github.com/munsheerck79/Ecom_project.git/pkg/api/handler"
	"github.com/munsheerck79/Ecom_project.git/pkg/config"
	"github.com/munsheerck79/Ecom_project.git/pkg/db"
	"github.com/munsheerck79/Ecom_project.git/pkg/repository"
	"github.com/munsheerck79/Ecom_project.git/pkg/usecase"
)

func InitiateAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnToDB,

		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,

		repository.NewAdminRepository,
		usecase.NewAdminService,
		handler.NewAdminHandler,

		repository.NewProductRepository,
		usecase.NewProductService,
		handler.NewProductHandler,

		repository.NewOrderRepository,
		usecase.NewOrderService,
		handler.NewOrderHandler,

		repository.NewPaymentRepository,
		usecase.NewPaymentService,
		handler.NewPaymentHandler,

		repository.NewCouponRepository,
		usecase.NewCouponService,
		handler.NewCouponHandler,

		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
