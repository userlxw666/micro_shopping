package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"micro_shopping/app/category/dao/model"
)

type Product struct {
	gorm.Model
	Name       string
	SKU        string
	Desc       string
	StockCount int
	Price      float32
	CategoryID uint
	Category   model.Category `json:"-"`
	IsDeleted  bool
}

func NewProduct(name string, desc string, stockCount int, price float32, cid uint) *Product {
	return &Product{
		Name:       name,
		SKU:        desc,
		StockCount: stockCount,
		Price:      price,
		CategoryID: cid,
		IsDeleted:  false,
	}
}

func (p *Product) CreateUUID() {
	Uuid := uuid.New().String()
	p.SKU = Uuid
}
