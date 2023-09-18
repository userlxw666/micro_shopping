package model

import (
	"gorm.io/gorm"
	"micro_shopping/app/user/dao/model"
)

type Order struct {
	gorm.Model
	User       model.User `gorm:"foreignKey:UserID"`
	UserID     uint
	OrderItem  []OrderItem `gorm:"foreignKey:OrderID"`
	TotalPrice float64
	IsCancel   bool
}

func NewOrder(userid uint, items []OrderItem) *Order {
	return &Order{
		UserID:    userid,
		OrderItem: items,
		IsCancel:  false,
	}
}
