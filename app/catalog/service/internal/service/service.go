package service

import (
	v1 "github.com/danielmoisa/ecommerce-microservices/api/catalog/service/v1"
	"github.com/danielmoisa/ecommerce-microservices/app/catalog/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCatalogService)

type CatalogService struct {
	v1.UnimplementedCatalogServer

	bc  *biz.ProductUseCase
	log *log.Helper
}

func NewCatalogService(bc *biz.ProductUseCase, logger log.Logger) *CatalogService {
	return &CatalogService{

		bc:  bc,
		log: log.NewHelper(log.With(logger, "module", "service/catalog"))}
}
