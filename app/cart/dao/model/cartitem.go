package model

import (
	"gorm.io/gorm"
	"micro_shopping/app/product/dao/model"
	"micro_shopping/idl/pb"
)

type CartItem struct {
	gorm.Model
	Product   model.Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Count     int
	Cart      Cart `gorm:"foreignKey:CartID"`
	CartID    uint
}

func NewCartItem(productID uint, count int, cartID uint) *CartItem {
	return &CartItem{
		ProductID: productID,
		Count:     count,
		CartID:    cartID,
	}
}

// 创建hooks 商品为0则删除商品
func (item *CartItem) AfterUpdate(tx *gorm.DB) error {
	if item.Count == 0 {
		return tx.Unscoped().Delete(&item).Error
	}
	return nil
}

func NewPbCartItem(items []*CartItem) []*pb.CartItem {
	var resp = make([]*pb.CartItem, 0)
	for _, v := range items {
		resp = append(resp, &pb.CartItem{
			ProductID: uint64(v.ProductID),
			Count:     int64(v.Count),
		})
	}
	return resp
}
