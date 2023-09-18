package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"micro_shopping/app/gateway/rpc"
	"micro_shopping/idl/pb"
	"micro_shopping/pkg/api_helper"
	"net/http"
)

func AddItem(c *gin.Context) {
	var req pb.AddItemRequest
	if err := c.ShouldBind(&req); err != nil {
		api_helper.HandleError(c, errors.New("绑定结构体失败"))
		return
	}

	UserId := c.Request.Context().Value("userid").(uint)
	req.UserID = uint64(UserId)
	_, err := rpc.AddItem(c, &req)
	if err != nil {
		api_helper.HandleError(c, errors.New("rpgc调用失败"))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "添加商品成功",
	})
}
func UpdateItem(c *gin.Context) {
	var req pb.UpdateItemRequest
	if err := c.ShouldBind(&req); err != nil {
		api_helper.HandleError(c, errors.New("绑定结构体失败"))
		return
	}

	UserId := c.Request.Context().Value("userid").(uint)
	req.UserID = uint64(UserId)
	_, err := rpc.UpdateItem(c, &req)
	if err != nil {
		api_helper.HandleError(c, errors.New("rpgc调用失败"))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "更新商品成功",
	})
}
func GetCart(c *gin.Context) {
	var req pb.GetCartRequest
	if err := c.ShouldBind(&req); err != nil {
		api_helper.HandleError(c, errors.New("绑定结构体失败"))
		return
	}

	UserId := c.Request.Context().Value("userid").(uint)
	req.UserID = uint64(UserId)
	resp, err := rpc.GetCart(c, &req)
	if err != nil {
		api_helper.HandleError(c, errors.New("rpgc调用失败"))
		return
	}
	c.JSON(http.StatusOK, resp.Items)
}
