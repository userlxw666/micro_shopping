package server

import (
	"context"
	"errors"
	"fmt"
	Orderdao "micro_shopping/app/order/dao"
	"micro_shopping/app/order/dao/model"
	"micro_shopping/idl/pb"
	"time"
)

type OrderService struct {
}

func (os *OrderService) CompleteOrder(ctx context.Context, req *pb.CompleteOrderRequest) (resp *pb.OrderEmpty, err error) {
	odDao := Orderdao.NewOrderDao(ctx)
	//odIDao := Orderdao.NewOrderItemDao(ctx)
	cartItemdao := Orderdao.NewCartItemDao(ctx)
	cartdao := Orderdao.NewCartDao(ctx)
	resp = new(pb.OrderEmpty)
	// 获取当前购物车
	currentCart, err := cartdao.CreateOrGetCartByID(uint(req.UserID))
	if err != nil {
		return nil, err
	}
	// 获取购物车中的商品
	cartItems, err := cartItemdao.GetAllCartItem(currentCart.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	orderitems := make([]model.OrderItem, 0)
	for _, item := range cartItems {
		orderitems = append(orderitems, *model.NewOrderItem(item.ProductID, item.Count))
	}
	err = odDao.CreateOrder(model.NewOrder(uint(req.UserID), orderitems))
	if err != nil {
		fmt.Println("2", err)
		return nil, err
	}
	return resp, nil
}
func (os *OrderService) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (resp *pb.OrderEmpty, err error) {
	resp = new(pb.OrderEmpty)
	odDao := Orderdao.NewOrderDao(ctx)
	// 获取当前订单
	order, err := odDao.GetOrderByID(uint(req.OrderID))
	if err != nil {
		return nil, err
	}
	// 判断是否超时
	if order.CreatedAt.Sub(time.Now()).Hours() > 336 {
		return nil, errors.New("订单已经超时")
	}
	order.IsCancel = true
	err = odDao.UpdateOrder(order)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (os *OrderService) GetOrders(ctx context.Context, req *pb.GetOrderRequest) (resp *pb.GetOrderResponse, err error) {
	odDao := Orderdao.NewOrderDao(ctx)
	orderitems, err := odDao.GetAllOrder(int(req.Pages.Page), int(req.Pages.PageSize), uint(req.UserID))
	resp = new(pb.GetOrderResponse)
	for _, order := range orderitems {
		for _, od := range order.OrderItem {
			resp.Items = append(resp.Items, &pb.OderItem{
				ProductID: uint64(od.ProductID),
				Count:     int64(od.Count),
			})
		}
	}
	return resp, nil
}
