package service

import (
	"context"

	v1 "github.com/danielmoisa/ecommerce-microservices/api/payment/service/v1"
)

func (s *PaymentService) PaymentAuth(ctx context.Context, req *v1.PaymentAuthReq) (reply *v1.PaymentAuthReply, err error) {
	return &v1.PaymentAuthReply{}, err
}
