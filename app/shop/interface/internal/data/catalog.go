package data

import (
	"context"
	"fmt"

	"golang.org/x/sync/singleflight"

	ctV1 "github.com/danielmoisa/ecommerce-microservices/api/catalog/service/v1"
	"github.com/danielmoisa/ecommerce-microservices/app/shop/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.CatalogRepo = (*catalogRepo)(nil)

type catalogRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func NewProductRepo(data *Data, logger log.Logger) biz.CatalogRepo {
	return &catalogRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/product")),
		sg:   &singleflight.Group{},
	}
}

func (r *catalogRepo) GetProduct(ctx context.Context, id int64) (*biz.Product, error) {
	result, err, _ := r.sg.Do(fmt.Sprintf("get_product_by_id_%d", id), func() (interface{}, error) {
		reply, err := r.data.bc.GetProduct(ctx, &ctV1.GetProductReq{
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
	})
	if err != nil {
		return nil, err
	}
	return result.(*biz.Product), nil
}

func (r *catalogRepo) ListProduct(ctx context.Context, pageNum, pageSize int64) ([]*biz.Product, error) {
	reply, err := r.data.bc.ListProduct(ctx, &ctV1.ListProductReq{
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
			Name:        x.Name,
			Description: x.Description,
			Count:       x.Count,
			Images:      images,
		})
	}
	return rv, err
}
