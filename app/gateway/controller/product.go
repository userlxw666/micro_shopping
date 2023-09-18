package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"micro_shopping/app/gateway/rpc"
	"micro_shopping/idl/pb"
	"micro_shopping/pkg/api_helper"
	"net/http"
)

func GetProducts(c *gin.Context) {
	var req pb.GetProductReq
	if err := c.ShouldBind(&req); err != nil {
		api_helper.HandleError(c, errors.New("结构体绑定失败"))
		return
	}
	resp, err := rpc.GetProducts(c, &req)
	if err != nil {
		api_helper.HandleError(c, errors.New("grpc调用失败"))
		return
	}

	c.JSON(http.StatusOK, resp.Resp)
}

func CreateProduct(c *gin.Context) {
	var req pb.CreateProductReq
	if err := c.ShouldBind(&req); err != nil {
		api_helper.HandleError(c, errors.New("结构体绑定失败"))
		return
	}
	_, err := rpc.CreateProduct(c, &req)
	if err != nil {
		api_helper.HandleError(c, errors.New("grpc调用失败"))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功！",
	})
}

func DeleteProduct(c *gin.Context) {
	var req pb.DeleteProductReq
	if err := c.ShouldBind(&req); err != nil {
		api_helper.HandleError(c, errors.New("结构体绑定失败"))
		return
	}
	_, err := rpc.DeleteProduct(c, &req)
	if err != nil {
		api_helper.HandleError(c, errors.New("grpc调用失败"))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功！",
	})
}

func UpdateProduct(c *gin.Context) {
	var req pb.UpdateProductReq
	if err := c.ShouldBind(&req); err != nil {
		api_helper.HandleError(c, errors.New("结构体绑定失败"))
		return
	}
	fmt.Println(req.SKU)
	_, err := rpc.UpdateProduct(c, &req)
	if err != nil {
		api_helper.HandleError(c, errors.New("grpc调用失败"))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功！",
	})
}
