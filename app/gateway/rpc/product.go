package rpc

import (
	"context"
	"fmt"
	"micro_shopping/idl/pb"
)

func GetProducts(ctx context.Context, req *pb.GetProductReq) (resp *pb.BulkProductResp, err error) {
	resp, err = ProductService.GetProducts(ctx, req)
	if err != nil {
		fmt.Println("获取grpc服务报错", err)
		return nil, err
	}
	return resp, nil
}

func CreateProduct(ctx context.Context, req *pb.CreateProductReq) (resp *pb.Empty, err error) {
	resp, err = ProductService.CreateProduct(ctx, req)
	if err != nil {
		fmt.Println("获取grpc服务报错", err)
		return nil, err
	}
	return resp, nil
}

func DeleteProduct(ctx context.Context, req *pb.DeleteProductReq) (resp *pb.Empty, err error) {
	resp, err = ProductService.DeleteProduct(ctx, req)
	if err != nil {
		fmt.Println("获取grpc服务报错", err)
		return nil, err
	}
	return resp, nil
}

func UpdateProduct(ctx context.Context, req *pb.UpdateProductReq) (resp *pb.Empty, err error) {
	resp, err = ProductService.UpdateProduct(ctx, req)
	if err != nil {
		fmt.Println("获取grpc服务报错", err)
		return nil, err
	}
	return resp, nil
}
