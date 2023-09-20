package Orderdao

import (
	"gorm.io/gorm"
	"micro_shopping/app/order/dao/model"
)

type OrderItemDao struct {
	*gorm.DB
}

//func NewOrderItemDao(ctx context.Context) *OrderItemDao {
//	if ctx == nil {
//		ctx = context.Background()
//	}
//	return &OrderItemDao{NewSQLClient(ctx)}
//}

// 创建订单项目
func (dao *OrderItemDao) CreateOrderItem(item *model.OrderItem) error {
	err := dao.Model(&model.OrderItem{}).Create(&item).Error
	return err
}

// 更新订单项目
func (dao *OrderItemDao) UpdateOrderItem(item *model.OrderItem) error {
	err := dao.Model(&model.OrderItem{}).Save(&item).Error
	return err
}
