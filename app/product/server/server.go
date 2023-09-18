package server

import (
	"context"
	"fmt"
	ProductDao "micro_shopping/app/product/dao"
	"micro_shopping/app/product/dao/model"
	"micro_shopping/idl/pb"
)

type ProductService struct {
}

// 获取商品
func (ps *ProductService) GetProducts(ctx context.Context, req *pb.GetProductReq) (resp *pb.BulkProductResp, err error) {
	resp = new(pb.BulkProductResp)
	pdDao := ProductDao.NewProductDao(ctx)
	if req.Text == "" {
		products, err := pdDao.GetAll(int(req.Page.Page), int(req.Page.PageSize))
		if err != nil {
			fmt.Println("get product error", err)
			return nil, err
		}
		result := resp.GetResp()
		for _, product := range products {
			result = append(result, BuildProduct(&product))
		}
		return resp, err
	} else {
		products, err := pdDao.SearchByString(req.Text, int(req.Page.Page), int(req.Page.PageSize))
		if err != nil {
			fmt.Println("search product error", err)
			return nil, err
		}
		var result []*pb.Product
		for _, product := range products {
			result = append(result, BuildProduct(&product))
		}
		resp.Resp = result
		return resp, err
	}
}

// 创建商品
func (ps *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductReq) (resp *pb.Empty, err error) {
	resp = new(pb.Empty)
	pdDao := ProductDao.NewProductDao(ctx)
	err = pdDao.CreateProduct(model.NewProduct(req.ProductName, req.Desc, int(req.StockCount), req.Price, uint(req.CategoryID)))
	if err != nil {
		fmt.Println("create product error", err)
		return nil, err
	}
	return resp, nil
}

// 删除商品
func (ps *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductReq) (resp *pb.Empty, err error) {
	resp = new(pb.Empty)
	pdDao := ProductDao.NewProductDao(ctx)
	err = pdDao.Delete(req.Sku)
	if err != nil {
		fmt.Println("delete product error", err)
		return nil, err
	}
	return resp, nil
}

// 更新商品
func (ps *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductReq) (resp *pb.Empty, err error) {
	resp = new(pb.Empty)
	pdDao := ProductDao.NewProductDao(ctx)
	pd := model.NewProduct(req.ProductName, req.Desc, int(req.StockCount), req.Price, uint(req.CategoryID))
	pd.SKU = req.SKU
	err = pdDao.UpdateProduct(pd)
	if err != nil {
		fmt.Println("update product error", err)
		return nil, err
	}
	return resp, nil
}

func BuildProduct(p *model.Product) *pb.Product {
	return &pb.Product{
		ProductName: p.Name,
		Desc:        p.Desc,
		Price:       p.Price,
		Sku:         p.SKU,
		CategoryID:  uint64(p.CategoryID),
		StockCount:  int64(p.StockCount),
	}
}
