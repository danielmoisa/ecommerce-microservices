package service

import (
	"context"

	v1 "github.com/danielmoisa/ecommerce-microservices/api/shop/admin/v1"
)

func (s *ShopAdmin) ListProduct(ctx context.Context, req *v1.ListProductReq) (*v1.ListProductReply, error) {
	rv, err := s.cc.ListProduct(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	reply := &v1.ListProductReply{
		Results: make([]*v1.ListProductReply_Product, 0),
	}
	for _, x := range rv {
		item := &v1.ListProductReply_Product{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
			Count:       x.Count,
			Image:       make([]*v1.ListProductReply_Product_Image, 0),
		}
		for _, img := range x.Images {
			item.Image = append(item.Image, &v1.ListProductReply_Product_Image{Url: img.URL})
		}
		reply.Results = append(reply.Results, item)
	}
	return reply, nil
}
func (s *ShopAdmin) GetProduct(ctx context.Context, req *v1.GetProductReq) (*v1.GetProductReply, error) {
	x, err := s.cc.GetProduct(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	reply := &v1.GetProductReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       make([]*v1.GetProductReply_Image, 0),
	}
	for _, img := range x.Images {
		reply.Image = append(reply.Image, &v1.GetProductReply_Image{Url: img.URL})
	}
	return reply, nil
}

func (s *ShopAdmin) CreateProduct(ctx context.Context, req *v1.CreateProductReq) (*v1.CreateProductReply, error) {
	return nil, nil
}

func (s *ShopAdmin) UpdateProduct(ctx context.Context, req *v1.UpdateProductReq) (*v1.UpdateProductReply, error) {
	return nil, nil
}

func (s *ShopAdmin) DeleteProduct(ctx context.Context, req *v1.DeleteProductReq) (*v1.DeleteProductReply, error) {
	return nil, nil
}
