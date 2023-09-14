package CategoryDao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"micro_shopping/app/category/dao/model"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &CategoryDao{NewSqlClient(ctx)}
}

func (dao *CategoryDao) CreateCategory(category *model.Category) {
	err := dao.Create(&category)
	if err != nil {
		fmt.Println("create category error", err)
		return
	}
}

func (dao *CategoryDao) FindByName(name string) []model.Category {
	var category []model.Category
	err := dao.Where("name=?", name).Find(&category)
	if err != nil {
		fmt.Println("find category error", err)
		return nil
	}
	return category
}

func (dao *CategoryDao) CreateAllCategory(categories []model.Category) (int, error) {
	var count int64
	err := dao.Create(&categories).Create(&count).Error
	return int(count), err
}

func (dao *CategoryDao) GetAll(pageIndex, pageSize int) ([]model.Category, int) {
	var categories []model.Category
	var count int64
	err := dao.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories).Count(&count).Error
	if err != nil {
		fmt.Println("get all categories error", err)
		return nil, 0
	}
	return categories, int(count)
}
