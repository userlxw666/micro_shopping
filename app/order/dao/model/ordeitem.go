package model

import (
	"errors"
	"gorm.io/gorm"
	"micro_shopping/app/product/dao/model"
)

type OrderItem struct {
	gorm.Model
	Product      model.Product `gorm:"foreignKey:ProductID"`
	ProductID    uint
	Count        int
	ProductPrice float32
	OrderID      uint
	IsCanceled   bool
}

func NewOrderItem(productId uint, count int, ProductPrice float32) *OrderItem {
	return &OrderItem{
		ProductID:    productId,
		Count:        count,
		ProductPrice: ProductPrice,
		IsCanceled:   false,
	}
}

// hooks 保存之前，更新产品库存
func (orderItem *OrderItem) BeforeCreate(tx *gorm.DB) error {
	var currentProduct model.Product
	var currentOrderItem OrderItem
	if err := tx.Where("id=?", orderItem.ProductID).First(&currentProduct).Error; err != nil {
		return err
	}

	reservedStockCount := 0
	if err := tx.Where("id=?", orderItem.ID).First(&currentOrderItem).Error; err != nil {
		reservedStockCount = currentOrderItem.Count
	}
	newStockCount := currentProduct.StockCount + reservedStockCount - orderItem.Count
	if newStockCount < 0 {
		return errors.New("库存不足")
	}
	if err := tx.Model(&currentProduct).Update("StockCount", newStockCount).Error; err != nil {
		return err
	}
	if orderItem.Count == 0 {
		err := tx.Unscoped().Delete(currentOrderItem).Error
		return err
	}
	return nil
}
