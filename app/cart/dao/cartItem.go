package Cartdao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"micro_shopping/app/cart/dao/model"
)

type CartItemDao struct {
	*gorm.DB
}

func NewCartItemDao(ctx context.Context) *CartItemDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &CartItemDao{NewSQLClient(ctx)}
}

// 创建cartItem
func (dao *CartItemDao) CreateCartItem(cartItem *model.CartItem) error {
	err := dao.Model(&model.CartItem{}).Create(&cartItem).Error
	return err
}

// 更新cartItem
func (dao *CartItemDao) UpdateCartItem(cartItem *model.CartItem) error {
	currentCartItem, err := dao.GetCartItemById(cartItem.ProductID, cartItem.CartID)
	if err != nil {
		fmt.Println("获取购物车项目失败", err)
		return err
	}
	err = dao.Model(&currentCartItem).Save(&cartItem).Error
	return err
}

// 通过productID和cartID获取cartItem
func (dao *CartItemDao) GetCartItemById(productId, cartId uint) (*model.CartItem, error) {
	var item *model.CartItem
	err := dao.Where(model.CartItem{ProductID: productId, CartID: cartId}).First(&item).Error
	if err != nil {
		return nil, err
	}
	return item, err
}

// 获取购物车中的所有商品
func (dao *CartItemDao) GetAllCartItem(cartId uint) ([]*model.CartItem, error) {
	var items []*model.CartItem
	err := dao.Where(&model.CartItem{CartID: cartId}).Find(&items).Error
	if err != nil {
		return nil, err
	}
	for i, item := range items {
		err = dao.Model(item).Association("Product").Find(&items[i].Product)
		if err != nil {
			return nil, err
		}
	}
	return items, nil
}
