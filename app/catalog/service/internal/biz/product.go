package biz

import (
	"context"
	"fmt"

	"github.com/danielmoisa/ecommerce-microservices/pkg/page_token"
	"golang.org/x/sync/singleflight"

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

type ProductRepo interface {
	CreateProduct(ctx context.Context, c *Product) (*Product, error)
	UpdateProduct(ctx context.Context, c *Product) (*Product, error)
	GetProduct(ctx context.Context, id int64) (*Product, error)
	ListProduct(ctx context.Context, pageNum, pageSize int64) ([]*Product, error)
	ListProductNext(ctx context.Context, start, end int32) ([]*Product, error)
	Count(ctx context.Context) (int, error)
}

type ProductUseCase struct {
	repo      ProductRepo
	log       *log.Helper
	pageToken page_token.ProcessPageTokens
	sg        singleflight.Group
}

func NewProductUseCase(repo ProductRepo, logger log.Logger) *ProductUseCase {
	return &ProductUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/product")), pageToken: page_token.NewTokenGenerate()}
}

func (uc *ProductUseCase) Create(ctx context.Context, u *Product) (*Product, error) {
	return uc.repo.CreateProduct(ctx, u)
}

func (uc *ProductUseCase) Get(ctx context.Context, id int64) (*Product, error) {
	return uc.repo.GetProduct(ctx, id)
}

func (uc *ProductUseCase) Update(ctx context.Context, u *Product) (*Product, error) {
	return uc.repo.UpdateProduct(ctx, u)
}

func (uc *ProductUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*Product, error) {
	return uc.repo.ListProduct(ctx, pageNum, pageSize)
}

func (uc *ProductUseCase) ListNext(ctx context.Context, pageSize int32, pageToken string) ([]*Product, string, error) {
	total, err := uc.repo.Count(ctx)
	if err != nil {
		return nil, "", err
	}

	start, end, nextToken, err := uc.pageToken.ProcessPageTokens(total, pageSize, pageToken)
	if err != nil {
		return nil, "", err
	}
	// single flight
	data, err, _ := uc.sg.Do(fmt.Sprintf("list_next_%d_%d", start, end), func() (interface{}, error) {
		return uc.repo.ListProductNext(ctx, start, end)
	})

	return data.([]*Product), nextToken, err
}
