package data

import (
	"context"

	"github.com/danielmoisa/ecommerce-microservices/app/shop/admin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	catalogv1 "github.com/danielmoisa/ecommerce-microservices/api/catalog/service/v1"
)

var _ biz.CatalogRepo = (*catalogRepo)(nil)

type catalogRepo struct {
	data *Data
	log  *log.Helper
}

func NewCatalogRepo(data *Data, logger log.Logger) biz.CatalogRepo {
	return &catalogRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/product")),
	}
}

func (r *catalogRepo) GetProduct(ctx context.Context, id int64) (*biz.Product, error) {
	reply, err := r.data.bc.GetProduct(ctx, &catalogv1.GetProductReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	images := make([]biz.Image, 0)
	for _, x := range reply.Image {
		images = append(images, biz.Image{URL: x.Url})
	}
	return &biz.Product{
		Id:          reply.Id,
		Name:        reply.Name,
		Description: reply.Description,
		Count:       reply.Count,
		Images:      images,
	}, err
}

func (r *catalogRepo) ListProduct(ctx context.Context, pageNum, pageSize int64) ([]*biz.Product, error) {
	reply, err := r.data.bc.ListProduct(ctx, &catalogv1.ListProductReq{
		PageNum:  pageNum,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Product, 0)
	for _, x := range reply.Results {
		images := make([]biz.Image, 0)
		for _, img := range x.Image {
			images = append(images, biz.Image{URL: img.Url})
		}
		rv = append(rv, &biz.Product{
			Id:          x.Id,
			Description: x.Description,
			Count:       x.Count,
			Images:      images,
		})
	}
	return rv, err
}

func (r *catalogRepo) CreateProduct(ctx context.Context, b *biz.Product) (*biz.Product, error) {
	images := make([]*catalogv1.CreateProductReq_Image, 0)
	for _, x := range b.Images {
		images = append(images, &catalogv1.CreateProductReq_Image{Url: x.URL})
	}
	reply, err := r.data.bc.CreateProduct(ctx, &catalogv1.CreateProductReq{
		Name:        b.Name,
		Description: b.Description,
		Count:       b.Count,
		Image:       images,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Product{
		Id: reply.Id,
	}, err
}

func (r *catalogRepo) UpdateProduct(ctx context.Context, b *biz.Product) (*biz.Product, error) {
	images := make([]*catalogv1.UpdateProductReq_Image, 0)
	for _, x := range b.Images {
		images = append(images, &catalogv1.UpdateProductReq_Image{Url: x.URL})
	}
	reply, err := r.data.bc.UpdateProduct(ctx, &catalogv1.UpdateProductReq{
		Name:        b.Name,
		Description: b.Description,
		Count:       b.Count,
		Image:       images,
	})
	if err != nil {
		return nil, err
	}

	return &biz.Product{
		Id: reply.Id,
	}, err
}
