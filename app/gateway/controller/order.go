package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"micro_shopping/app/gateway/rpc"
	"micro_shopping/idl/pb"
	"micro_shopping/pkg/api_helper"
	"net/http"
)

func CompleteOrder(c *gin.Context) {
	var req pb.CompleteOrderRequest
	UserID := c.Request.Context().Value("userid").(uint)
	req.UserID = uint64(UserID)
	if _, err := rpc.CompleteOrder(c, &req); err != nil {
		api_helper.HandleError(c, errors.New("rpc调用失败"))
		return
	}
	c.JSON(http.StatusOK, api_helper.Response{Msg: "完成订单"})
}
func CancelOrder(c *gin.Context) {
	var req pb.CancelOrderRequest
	if err := c.ShouldBind(&req); err != nil {
		api_helper.HandleError(c, errors.New("结构体绑定失败"))
		return
	}
	UserID := c.Request.Context().Value("userid").(uint)
	req.UserID = uint64(UserID)
	if _, err := rpc.CancelOrder(c, &req); err != nil {
		api_helper.HandleError(c, errors.New("rpc调用失败"))
		return
	}
	c.JSON(http.StatusOK, api_helper.Response{Msg: "订单取消成功"})
}
func GetOrders(c *gin.Context) {
	var req pb.GetOrderRequest
	if err := c.ShouldBind(&req.Pages); err != nil {
		api_helper.HandleError(c, errors.New("结构体绑定失败"))
		return
	}
	UserID := c.Request.Context().Value("userid").(uint)
	req.UserID = uint64(UserID)
	resp, err := rpc.GetOrders(c, &req)
	if err != nil {
		api_helper.HandleError(c, errors.New("rpc调用失败"))
		return
	}
	c.JSON(http.StatusOK, api_helper.Response{Msg: resp})
}
