package rpc

import (
	"context"
	"fmt"
	"micro_shopping/idl/pb"
)

func CompleteOrder(ctx context.Context, req *pb.CompleteOrderRequest) (resp *pb.OrderEmpty, err error) {
	resp, err = OrderService.CompleteOrder(ctx, req)
	if err != nil {
		fmt.Println("获取grpc服务报错", err)
		return nil, err
	}
	return resp, nil
}

func CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (resp *pb.OrderEmpty, err error) {
	resp, err = OrderService.CancelOrder(ctx, req)
	if err != nil {
		fmt.Println("获取grpc服务报错", err)
		return nil, err
	}
	return resp, nil
}

func GetOrders(ctx context.Context, req *pb.GetOrderRequest) (resp *pb.GetOrderResponse, err error) {
	resp, err = OrderService.GetOrders(ctx, req)
	if err != nil {
		fmt.Println("获取grpc服务报错", err)
		return nil, err
	}
	return resp, nil
}
