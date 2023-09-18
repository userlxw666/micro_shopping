package server

import (
	"context"
	"errors"
	"fmt"
	"micro_shopping/app/cart/dao"
	"micro_shopping/app/cart/dao/model"
	"micro_shopping/idl/pb"
)

type CartService struct {
}

func (cs *CartService) AddItem(ctx context.Context, req *pb.AddItemRequest) (resp *pb.CartEmpty, err error) {
	resp = new(pb.CartEmpty)
	PtDao := Cartdao.NewProduct(ctx)
	CtDao := Cartdao.NewCartDao(ctx)
	CtItDao := Cartdao.NewCartItemDao(ctx)
	//获取当前商品
	currentProduct := PtDao.SearchBySKU(req.Sku)
	//获取当前购物车,没有就创建一个购物车
	currentCart, err := CtDao.CreateOrGetCartByID(uint(req.UserID))
	if err != nil {
		fmt.Println("获取购物车或创建购物车失败", err)
		return nil, err
	}

	if req.Count <= 0 {
		return resp, errors.New("错误数量")
	}
	currentItem, err := CtItDao.GetCartItemById(currentProduct.ID, currentCart.ID)
	if err != nil {
		// 直接创建购物车项
		if currentProduct.StockCount < int(req.Count) {
			fmt.Println("库存不足")
			return nil, errors.New("库存不足")
		}
		err = CtItDao.CreateCartItem(model.NewCartItem(currentProduct.ID, int(req.Count), currentCart.ID))
		if err != nil {
			fmt.Println("creat cartitem error", err)
			return nil, err
		}
	}

	if currentProduct.StockCount-currentItem.Count < int(req.Count) {
		return resp, errors.New("库存不够")
	}

	// 更新所添加的数量
	currentItem.Count += int(req.Count)
	err = CtItDao.UpdateCartItem(currentItem)
	if err == nil {
		currentProduct.StockCount -= int(req.Count)
		err = PtDao.UpdateProduct(currentProduct)
		if err != nil {
			fmt.Println("商品数量更新失败", err)
			return resp, err
		}
	}
	return resp, err
}

func (cs *CartService) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (resp *pb.CartEmpty, err error) {
	resp = new(pb.CartEmpty)
	PtDao := Cartdao.NewProduct(ctx)
	CtDao := Cartdao.NewCartDao(ctx)
	CtItDao := Cartdao.NewCartItemDao(ctx)
	//获取当前商品
	currentProduct := PtDao.SearchBySKU(req.Sku)
	//获取当前购物车
	currentCart, err := CtDao.GetCartByID(uint(req.UserID))
	if err != nil {
		fmt.Println("获取购物车失败", err)
		return nil, err
	}
	currentItem, err := CtItDao.GetCartItemById(uint(currentProduct.ID), uint(currentCart.ID))
	if err != nil {
		fmt.Println("获取cartItem失败", err)
		return nil, err
	}
	if currentProduct.StockCount+currentItem.Count < int(req.Count) {
		return nil, errors.New("库存不够")
	}
	currentItem.Count = int(req.Count)
	err = CtItDao.UpdateCartItem(currentItem)
	if err == nil {
		currentProduct.StockCount += currentItem.Count
		currentProduct.StockCount -= int(req.Count)
		err = PtDao.UpdateProduct(currentProduct)
		if err != nil {
			fmt.Println("商品数量更新失败", err)
			return nil, err
		}
	}
	return resp, nil
}

func (cs *CartService) GetCart(ctx context.Context, req *pb.GetCartRequest) (resp *pb.GetCartResponse, err error) {
	CtDao := Cartdao.NewCartDao(ctx)
	CtItDao := Cartdao.NewCartItemDao(ctx)
	resp = new(pb.GetCartResponse)
	currentCart, err := CtDao.CreateOrGetCartByID(uint(req.UserID))
	if err != nil {
		fmt.Println("获取购物车或创建购物车失败", err)
		return resp, err
	}
	items, err := CtItDao.GetAllCartItem(currentCart.ID)
	if err != nil {
		fmt.Println("获取购物车项目失败", err)
		return resp, err
	}
	resp = BuilerResp(int64(req.UserID), model.NewPbCartItem(items))
	return resp, nil
}

func BuilerResp(id int64, items []*pb.CartItem) *pb.GetCartResponse {
	return &pb.GetCartResponse{
		UserID: id,
		Items:  items,
	}
}
