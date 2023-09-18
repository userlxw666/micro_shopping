package server

import (
	"context"
	"errors"
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
