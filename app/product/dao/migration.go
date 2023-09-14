package ProductDao

import (
	"fmt"
	"micro_shopping/app/product/dao/model"
)

func Migration() {
	err := DB.AutoMigrate(&model.Product{})
	if err != nil {
		fmt.Println("create product table error", err)
		return
	}
}
