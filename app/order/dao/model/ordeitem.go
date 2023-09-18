package model

import (
	"gorm.io/gorm"
	"micro_shopping/app/product/dao/model"
)

type OrderItem struct {
	gorm.Model
	Product    model.Product `gorm:"foreignKey:ProductID"`
	ProductID  uint
	Count      int
	OrderID    uint
	IsCanceled bool
}

func NewOrderItem(productId uint, count int) *OrderItem {
	return &OrderItem{
		ProductID:  productId,
		Count:      count,
		IsCanceled: false,
	}
}
