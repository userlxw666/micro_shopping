package Orderdao

import (
	"fmt"
	"micro_shopping/app/order/dao/model"
)

func MigrationOrderItem() {
	err := DB.AutoMigrate(&model.OrderItem{})
	if err != nil {
		fmt.Println("create orderitem table error", err)
		return
	}
}

func MigrationOrder() {
	err := DB.AutoMigrate(&model.Orders{})
	if err != nil {
		fmt.Println("create order table error", err)
		return
	}
}
