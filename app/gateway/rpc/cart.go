package rpc

import (
	"context"
	"fmt"
	"micro_shopping/idl/pb"
)

func AddItem(ctx context.Context, res *pb.AddItemRequest) (resp *pb.CartEmpty, err error) {
	resp, err = CartService.AddItem(ctx, res)
	if err != nil {
		fmt.Println("获取grpc服务报错", err)
		return nil, err
	}
	return resp, nil
}
func UpdateItem(ctx context.Context, res *pb.UpdateItemRequest) (resp *pb.CartEmpty, err error) {
	resp, err = CartService.UpdateItem(ctx, res)
	if err != nil {
		fmt.Println("获取grpc服务报错", err)
		return nil, err
	}
	return resp, nil
}
func GetCart(ctx context.Context, res *pb.GetCartRequest) (resp *pb.GetCartResponse, err error) {
	resp, err = CartService.GetCart(ctx, res)
	if err != nil {
		fmt.Println("获取grpc服务报错", err)
		return nil, err
	}
	return resp, nil
}
