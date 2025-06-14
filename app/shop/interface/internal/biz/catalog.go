package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Image struct {
	URL string
}

type Product struct {
	Id          int64
	Name        string
	Description string
	Count       int64
	Images      []Image
}

type CatalogRepo interface {
	GetProduct(ctx context.Context, id int64) (*Product, error)
	ListProduct(ctx context.Context, pageNum, pageSize int64) ([]*Product, error)
}

type CatalogUseCase struct {
	repo CatalogRepo
	log  *log.Helper
}

func NewCatalogUseCase(repo CatalogRepo, logger log.Logger) *CatalogUseCase {
	return &CatalogUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/product"))}
}

func (uc *CatalogUseCase) GetProduct(ctx context.Context, id int64) (*Product, error) {
	return uc.repo.GetProduct(ctx, id)
}

func (uc *CatalogUseCase) ListProduct(ctx context.Context, pageNum, pageSize int64) ([]*Product, error) {
	return uc.repo.ListProduct(ctx, pageNum, pageSize)
}
