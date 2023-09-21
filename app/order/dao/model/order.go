package model

import (
	"fmt"
	"gorm.io/gorm"
	"micro_shopping/app/cart/dao/model"
	ProductModel "micro_shopping/app/product/dao/model"
	UserModel "micro_shopping/app/user/dao/model"
)

type Orders struct {
	gorm.Model
	User       UserModel.User `gorm:"foreignKey:UserID"`
	UserID     uint
	OrderItem  []OrderItem `gorm:"foreignKey:OrderID"`
	TotalPrice float64
	IsCancel   bool
}

func NewOrder(userid uint, items []OrderItem) *Orders {
	var total float32 = 0.0
	for _, item := range items {
		total += item.ProductPrice * float32(item.Count)
		fmt.Println(item.Product.Price, float32(item.Count))
	}
	return &Orders{
		UserID:     userid,
		OrderItem:  items,
		IsCancel:   false,
		TotalPrice: float64(total),
	}
}

// hooks 创建order前查找购物车并删除购物车
func (order *Orders) BeforeCreate(tx *gorm.DB) error {
	var currentCart model.Cart
	// 通过用户id查找购物车
	if err := tx.Where("user_id=?", order.UserID).First(&currentCart).Error; err != nil {
		return err
	}
	// 删除购物车项目
	if err := tx.Where("cart_id=?", currentCart.ID).Unscoped().Delete(&model.CartItem{}).Error; err != nil {
		return err
	}
	// 删除购物车
	if err := tx.Unscoped().Delete(&currentCart).Error; err != nil {
		return err
	}
	return nil
}

// 如果订单被取消，数量将返回产品库存
func (order *Orders) BeforeUpdate(tx *gorm.DB) error {
	if order.IsCancel {
		var orderedItems []OrderItem
		// 寻找订单项目
		if err := tx.Where("order_id=?", order.ID).Find(&orderedItems).Error; err != nil {
			return err
		}
		// 对订单项遍历，将订单中的商品数量返回到商品库存中
		for _, item := range orderedItems {
			var currentProduct ProductModel.Product
			// 找到当前商品
			if err := tx.Where("id=?", item.ProductID).First(&currentProduct).Error; err != nil {
				return err
			}
			// 将订单商品数量返回到商品库存
			newStockCount := currentProduct.StockCount + item.Count
			if err := tx.Model(&currentProduct).Update("stock_count", newStockCount).Error; err != nil {
				return err
			}
			// 订单项设置为取消
			if err := tx.Model(&item).Update("is_canceled", true).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
