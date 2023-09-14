package rpc

import (
	"context"
	"fmt"
	"micro_shopping/idl/pb"
)

func CreateCategory(ctx context.Context, req *pb.CategoryRequest) (resp *pb.CategoryResponse, err error) {
	resp, err = CategoryService.CreateCategory(ctx, req)
	if err != nil {
		fmt.Println("grpc服务报错", err)
		return nil, err
	}
	return
}

func BulkCreateCategory(ctx context.Context, req *pb.BulkRequest) (resp *pb.BulkResponse, err error) {
	resp, err = CategoryService.BulkCreateCategory(ctx, req)
	if err != nil {
		fmt.Println("grpc服务报错", err)
		return nil, err
	}
	return
}

func GetCategories(ctx context.Context, page *pb.Page) (resp *pb.BulkResponse, err error) {
	resp, err = CategoryService.GetCategories(ctx, page)
	if err != nil {
		fmt.Println("grpc服务报错", err)
		return nil, err
	}
	return
}
