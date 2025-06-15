package data

import (
	"context"

	"github.com/danielmoisa/ecommerce-microservices/app/catalog/service/internal/biz"
	"github.com/danielmoisa/ecommerce-microservices/pkg/util/pagination"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.ProductRepo = (*productRepo)(nil)

type productRepo struct {
	data *Data
	log  *log.Helper
}

func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/product")),
	}
}

func (r *productRepo) CreateProduct(ctx context.Context, b *biz.Product) (*biz.Product, error) {
	po, err := r.data.db.Product.
		Create().
		SetName(b.Name).
		SetDescription(b.Description).
		SetCount(b.Count).
		SetImages(b.Images).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.Product{
		Id:          po.ID,
		Description: po.Description,
		Count:       po.Count,
		Images:      po.Images,
	}, nil
}

func (r *productRepo) GetProduct(ctx context.Context, id int64) (*biz.Product, error) {
	po, err := r.data.db.Product.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Product{
		Id:          po.ID,
		Description: po.Description,
		Count:       po.Count,
		Images:      po.Images,
	}, nil
}

func (r *productRepo) UpdateProduct(ctx context.Context, b *biz.Product) (*biz.Product, error) {
	po, err := r.data.db.Product.
		Create().
		SetName(b.Name).
		SetDescription(b.Description).
		SetCount(b.Count).
		SetImages(b.Images).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.Product{
		Id:          po.ID,
		Description: po.Description,
		Count:       po.Count,
		Images:      po.Images,
	}, nil
}

func (r *productRepo) ListProduct(ctx context.Context, pageNum, pageSize int64) ([]*biz.Product, error) {
	pos, err := r.data.db.Product.Query().
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Limit(int(pageSize)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Product, 0, len(pos))
	for _, po := range pos {
		rv = append(rv, &biz.Product{
			Id:          po.ID,
			Name:        po.Name,
			Description: po.Description,
			Count:       po.Count,
			Images:      po.Images,
		})
	}
	return rv, nil
}

func (r *productRepo) Count(ctx context.Context) (int, error) {
	return r.data.db.Product.Query().Count(ctx)
}

func (r *productRepo) ListProductNext(ctx context.Context, start, end int32) ([]*biz.Product, error) {

	pos, err := r.data.db.Product.Query().
		Offset(int(start)).
		Limit(int(end - start)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Product, 0, len(pos))
	for _, po := range pos {
		rv = append(rv, &biz.Product{
			Id:          po.ID,
			Description: po.Description,
			Count:       po.Count,
			Images:      po.Images,
		})
	}
	return rv, nil
}
