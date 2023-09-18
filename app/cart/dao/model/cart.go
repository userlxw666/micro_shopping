package model

import (
	"gorm.io/gorm"
	"micro_shopping/app/user/dao/model"
)

type Cart struct {
	gorm.Model
	User   model.User `gorm:"foreignKey:UserID"`
	UserID uint
}

func NewCart(userId uint) *Cart {
	return &Cart{
		UserID: userId,
	}
}
