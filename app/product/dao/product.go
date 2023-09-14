package ProductDao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"micro_shopping/app/product/dao/model"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &ProductDao{NewSqlClient(ctx)}
}

// 创建商品
func (p *ProductDao) CreateProduct(product *model.Product) error {
	result := p.Create(p)
	return result.Error
}

// 更新
func (p *ProductDao) UpdateProduct(product *model.Product) error {
	return p.Save(&product).Error
}

// 通过关键字寻找商品
func (p *ProductDao) SearchByString(str string, pageIndex, pageSize int) ([]model.Product, int) {
	var products []model.Product
	var count int64
	converStr := "%" + str + "%"
	err := p.Where("name like ? or sku like ?", converStr, converStr).Where(
		"IsDeleted=?", false).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count).Error
	if err != nil {
		fmt.Println("search error", err)
		return nil, 0
	}
	return products, int(count)
}

// 通过sku寻找商品
func (p *ProductDao) SearchBySKU(sku string) *model.Product {
	var product model.Product
	err := p.Where("IsDeleted=?", false).Where(model.Product{SKU: sku}).First(&product).Error
	if err != nil {
		fmt.Println("search product bu sku error", err)
		return nil
	}
	return &product
}

// 查询所有商品
func (p *ProductDao) GetAll(page, pageSize int) ([]model.Product, int) {
	var products []model.Product
	var count int64
	err := p.Offset((page - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)
	if err != nil {
		fmt.Println("获取所有商品失败", err)
		return nil, 0
	}
	return products, int(count)
}

// 删除商品
func (p *ProductDao) Delete(sku string) error {
	product := p.SearchBySKU(sku)
	product.IsDeleted = true

	err := p.Save(&product).Error
	return err
}
