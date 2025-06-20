// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/danielmoisa/ecommerce-microservices/app/shop/admin/internal/biz"
	"github.com/danielmoisa/ecommerce-microservices/app/shop/admin/internal/conf"
	"github.com/danielmoisa/ecommerce-microservices/app/shop/admin/internal/data"
	"github.com/danielmoisa/ecommerce-microservices/app/shop/admin/internal/server"
	"github.com/danielmoisa/ecommerce-microservices/app/shop/admin/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	userClient := data.NewUserServiceClient(discovery, tracerProvider)
	cartClient := data.NewCartServiceClient(discovery, tracerProvider)
	catalogClient := data.NewCatalogServiceClient(discovery, tracerProvider)
	dataData, err := data.NewData(confData, logger, userClient, cartClient, catalogClient)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger, userClient)
	catalogRepo := data.NewCatalogRepo(dataData, logger)
	catalogUseCase := biz.NewCatalogUseCase(catalogRepo, logger)
	shopAdmin := service.NewShopAdmin(userUseCase, catalogUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, logger, tracerProvider, shopAdmin)
	grpcServer := server.NewGRPCServer(confServer, logger, tracerProvider, shopAdmin)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
	}, nil
}
