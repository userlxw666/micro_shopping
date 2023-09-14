package server

import (
	"context"
	"errors"
	"fmt"
	"micro_shopping/app/category/dao"
	"micro_shopping/app/category/dao/model"
	"micro_shopping/idl/pb"
)

type CategoryService struct {
}

func (cgsv *CategoryService) CreateCategory(ctx context.Context, req *pb.CategoryRequest) (resp *pb.CategoryResponse, err error) {
	resp = new(pb.CategoryResponse)
	categoryDao := CategoryDao.NewCategoryDao(ctx)
	categories := categoryDao.FindByName(req.Name)
	if len(categories) > 0 {
		return resp, errors.New("name existed")
	}
	category := model.NewCategory(req.Name, req.Desc)
	categoryDao.CreateCategory(category)
	resp.Name = req.Name
	resp.Desc = req.Desc
	return resp, nil
}

func (cgsv *CategoryService) BulkCreateCategory(ctx context.Context, req *pb.BulkRequest) (resp *pb.BulkResponse, err error) {
	resp = new(pb.BulkResponse)
	result := req.GetBulkRequest()
	categoryDao := CategoryDao.NewCategoryDao(ctx)
	var categories = make([]model.Category, 0)
	rp := resp.GetBulkResponse()
	// 判断该种类是否已经存在
	for _, request := range result {
		cat := categoryDao.FindByName(request.Name)
		if cat != nil {
			fmt.Println("categories exist", err)
			return resp, err
		}
		categories = append(categories, *model.NewCategory(request.Name, request.Desc))
		rp = append(rp, &pb.CategoryResponse{
			Name: request.Name,
			Desc: request.Desc,
		})
	}
	// 存入数据库
	_, err = categoryDao.CreateAllCategory(categories)
	if err != nil {
		fmt.Println("批量创建失败", err)
		return nil, err
	}
	return resp, nil
}

func (cgsv *CategoryService) GetCategories(ctx context.Context, page *pb.Page) (resp *pb.BulkResponse, err error) {
	resp = new(pb.BulkResponse)
	categoryDao := CategoryDao.NewCategoryDao(ctx)
	categories, _ := categoryDao.GetAll(int(page.Page), int(page.PageSize))
	var a = make([]*pb.CategoryResponse, len(categories))
	for k, category := range categories {
		a[k] = &pb.CategoryResponse{
			Name: category.Name,
			Desc: category.Desc,
		}
		resp.BulkResponse = append(resp.BulkResponse, a[k])
	}
	return resp, nil
}
