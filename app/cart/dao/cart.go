package Cartdao

import (
	"context"
	"gorm.io/gorm"
	"micro_shopping/app/cart/dao/model"
	ProductDao "micro_shopping/app/product/dao"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &CartDao{NewSQLClient(ctx)}
}

func NewProduct(ctx context.Context) *ProductDao.ProductDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &ProductDao.ProductDao{NewSQLClient(ctx)}
}

// 创建购物车
func (dao *CartDao) CreateCart(cart *model.Cart) error {
	err := dao.Model(&model.Cart{}).Create(&cart).Error
	return err
}

// 更新购物车
func (dao *CartDao) UpDateCart(cart *model.Cart) error {
	err := dao.Model(&model.Cart{}).Save(&cart).Error
	return err
}

// 根据用户id查找或创建购物车
func (dao *CartDao) CreateOrGetCartByID(userId uint) (*model.Cart, error) {
	var cart *model.Cart
	err := dao.Where(&model.Cart{UserID: userId}).Attrs(model.NewCart(userId)).FirstOrCreate(&cart).Error
	return cart, err
}

// 根据用户id查找购物车
func (dao *CartDao) GetCartByID(userId uint) (*model.Cart, error) {
	var cart *model.Cart
	err := dao.Where(&model.Cart{UserID: userId}).First(&cart).Error
	return cart, err
}
