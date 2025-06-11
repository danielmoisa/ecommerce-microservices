package service

import (
	"context"

	v1 "github.com/danielmoisa/ecommerce-microservices/api/shop/admin/v1"
)

func (s *ShopAdmin) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	return nil, nil
}

func (s *ShopAdmin) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	return nil, nil
}
