package server

import (
	"context"
	"micro_shopping/idl/pb"
)

type OrderService struct {
}

func (os *OrderService) CompleteOrder(ctx context.Context, req *pb.CompleteOrderRequest) (resp *pb.Empty, err error) {
	return
}
func (os *OrderService) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (resp *pb.Empty, err error) {
	return
}
func (os *OrderService) GetOrders(ctx context.Context, req *pb.GetOrderRequest) (resp *pb.GetOrderResponse, err error) {
	return
}
