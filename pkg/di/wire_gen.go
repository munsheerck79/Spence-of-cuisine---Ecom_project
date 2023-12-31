// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/munsheerck79/Ecom_project.git/pkg/api"
	"github.com/munsheerck79/Ecom_project.git/pkg/api/handler"
	"github.com/munsheerck79/Ecom_project.git/pkg/config"
	"github.com/munsheerck79/Ecom_project.git/pkg/db"
	"github.com/munsheerck79/Ecom_project.git/pkg/repository"
	"github.com/munsheerck79/Ecom_project.git/pkg/usecase"
)

// Injectors from wire.go:

func InitiateAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnToDB(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userService := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userService)
	adminRepository := repository.NewAdminRepository(gormDB)
	adminService := usecase.NewAdminService(adminRepository)
	adminHandler := handler.NewAdminHandler(adminService)
	productRepository := repository.NewProductRepository(gormDB)
	productService := usecase.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	orderRepository := repository.NewOrderRepository(gormDB)
	paymentRepository := repository.NewPaymentRepository(gormDB)
	orderService := usecase.NewOrderService(orderRepository, productRepository, userRepository, paymentRepository, userService)
	orderHandler := handler.NewOrderHandler(orderService, userService)
	paymentService := usecase.NewPaymentService(paymentRepository, orderRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService, orderService, userService)
	couponRepository := repository.NewCouponRepository(gormDB)
	couponService := usecase.NewCouponService(couponRepository)
	couponHandler := handler.NewCouponHandler(couponService)
	serverHTTP := http.NewServerHTTP(userHandler, adminHandler, productHandler, orderHandler, paymentHandler, couponHandler)
	return serverHTTP, nil
}
