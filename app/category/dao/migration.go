package CategoryDao

import (
	"fmt"
	"micro_shopping/app/category/dao/model"
)

func Migration() {
	err := DB.AutoMigrate(&model.Category{})
	if err != nil {
		fmt.Println("create category table error", err)
		return
	}
}
