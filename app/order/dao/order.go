package Orderdao

import (
	"context"
	"gorm.io/gorm"
	Cartdao "micro_shopping/app/cart/dao"
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

func NewCartDao(ctx context.Context) *Cartdao.CartDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Cartdao.CartDao{NewSQLClient(ctx)}
}

func NewCartItemDao(ctx context.Context) *Cartdao.CartItemDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Cartdao.CartItemDao{NewSQLClient(ctx)}
}

// 创建订单
func (dao *OrderDao) CreateOrder(order *model.Orders) error {
	err := dao.Model(&model.Orders{}).Create(&order).Error
	return err
}

// 根据订单id查找订单
func (dao *OrderDao) GetOrderByID(oid uint) (*model.Orders, error) {
	var order *model.Orders
	err := dao.Where("id=?", oid).Where("is_cancel=?", false).First(&order).Error
	return order, err
}

// 查找所有订单
func (dao *OrderDao) GetAllOrder(pageIndex, pageSize int, uid uint) ([]model.Orders, error) {
	var orders []model.Orders

	err := dao.Where("is_cancel=?", false).Where("user_id=?", uid).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&orders).Error
	for k, order := range orders {
		dao.Where("order_id=?", order.ID).Find(&orders[k].OrderItem)
		for i, product := range orders[k].OrderItem {
			dao.Where("id=?", product.ID).First(&orders[k].OrderItem[i].Product)
		}
	}
	return orders, err
}

// 更新订单
func (dao *OrderDao) UpdateOrder(order *model.Orders) error {
	err := dao.Save(&order).Error
	return err
}
