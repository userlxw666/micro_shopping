package Cartdao

import (
	"fmt"
	"micro_shopping/app/cart/dao/model"
)

func MigrationCart() {
	err := DB.AutoMigrate(&model.Cart{})
	if err != nil {
		fmt.Println("create cart table error", err)
		return
	}
}

func MigrationCartItem() {
	err := DB.AutoMigrate(&model.CartItem{})
	if err != nil {
		fmt.Println("create cartItem error", err)
		return
	}
}
