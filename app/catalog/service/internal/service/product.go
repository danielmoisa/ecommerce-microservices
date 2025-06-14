package service

import (
	"context"

	v1 "github.com/danielmoisa/ecommerce-microservices/api/catalog/service/v1"
	"github.com/danielmoisa/ecommerce-microservices/app/catalog/service/internal/biz"
)

func (s *CatalogService) CreateProduct(ctx context.Context, req *v1.CreateProductReq) (*v1.CreateProductReply, error) {
	b := &biz.Product{
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
		Images:      make([]biz.Image, 0),
	}
	for _, x := range req.Image {
		b.Images = append(b.Images, biz.Image{URL: x.Url})
	}
	x, err := s.bc.Create(ctx, b)
	img := make([]*v1.CreateProductReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1.CreateProductReply_Image{Url: i.URL})
	}
	return &v1.CreateProductReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) GetProduct(ctx context.Context, req *v1.GetProductReq) (*v1.GetProductReply, error) {
	x, err := s.bc.Get(ctx, req.Id)
	img := make([]*v1.GetProductReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1.GetProductReply_Image{Url: i.URL})
	}
	return &v1.GetProductReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) UpdateProduct(ctx context.Context, req *v1.UpdateProductReq) (*v1.UpdateProductReply, error) {
	b := &biz.Product{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
		Images:      make([]biz.Image, 0),
	}
	for _, x := range req.Image {
		b.Images = append(b.Images, biz.Image{URL: x.Url})
	}
	x, err := s.bc.Update(ctx, b)
	img := make([]*v1.UpdateProductReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1.UpdateProductReply_Image{Url: i.URL})
	}
	return &v1.UpdateProductReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) ListProduct(ctx context.Context, req *v1.ListProductReq) (*v1.ListProductReply, error) {
	rv, err := s.bc.List(ctx, req.PageNum, req.PageSize)
	rs := make([]*v1.ListProductReply_Product, 0)
	for _, x := range rv {
		img := make([]*v1.ListProductReply_Product_Image, 0)
		for _, i := range x.Images {
			img = append(img, &v1.ListProductReply_Product_Image{Url: i.URL})
		}
		rs = append(rs, &v1.ListProductReply_Product{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
			Image:       img,
		})
	}
	return &v1.ListProductReply{
		Results: rs,
	}, err
}

func (s *CatalogService) ListProductNextToken(ctx context.Context, req *v1.ListProductNextTokenReq) (*v1.ListProductReplyNextToken, error) {
	rv, token, err := s.bc.ListNext(ctx, req.PageSize, req.PageToken)
	rs := make([]*v1.ListProductReplyNextToken_Product, 0)
	for _, x := range rv {
		img := make([]*v1.ListProductReplyNextToken_Product_Image, 0)
		for _, i := range x.Images {
			img = append(img, &v1.ListProductReplyNextToken_Product_Image{Url: i.URL})
		}
		rs = append(rs, &v1.ListProductReplyNextToken_Product{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
			Image:       img,
		})
	}
	return &v1.ListProductReplyNextToken{
		Results:       rs,
		NextPageToken: token,
	}, err
}
