package dao

import (
	"fmt"
	"micro_shopping/app/user/dao/model"
)

func Migration() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("create user table error", err)
		return
	}
}
