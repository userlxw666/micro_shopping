package Orderdao

import (
	"context"
	"gorm.io/gorm"
	"micro_shopping/app/order/dao/model"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &OrderDao{NewSQLClient(ctx)}
}

// 创建订单
func (dao *OrderDao) CreateOrder(order *model.Order) error {
	err := dao.Model(&model.Order{}).Create(*order).Error
	return err
}

// 根据订单id查找订单
func (dao *OrderDao) GetOrderByID(oid uint) (*model.Order, error) {
	var order *model.Order
	err := dao.Where("ID=?", oid).Where("IsCanceled=?", false).First(&order).Error
	return order, err
}

// 查找所有订单
func (dao *OrderDao) GetAllOrder(pageIndex, pageSize int, uid uint) ([]model.Order, error) {
	var orders []model.Order

	err := dao.Where("IsCanceled=?", false).Where("UserID=?", uid).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&orders).Error
	for k, order := range orders {
		dao.Where("OrderID=?", order.ID).Find(orders[k].OrderItem)
		for i, product := range orders[k].OrderItem {
			dao.Where("ID=?", product.ID).First(orders[k].OrderItem[i].Product)
		}
	}
	return orders, err
}

// 更新订单
func (dao *OrderDao) UpdateOrder(order *model.Order) error {
	err := dao.Model(&model.Order{}).Save(&order).Error
	return err
}
