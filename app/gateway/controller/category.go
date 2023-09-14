package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"micro_shopping/app/gateway/rpc"
	"micro_shopping/idl/pb"
	"micro_shopping/pkg/api_helper"
	"micro_shopping/pkg/utils"
	"net/http"
)

func CreateCategory(c *gin.Context) {
	var req pb.CategoryRequest
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println("结构体绑定失败", err)
		return
	}

	// 调用rpc
	resp, err := rpc.CreateCategory(c, &req)
	if err != nil {
		api_helper.HandleError(c, errors.New("category rpc 调用失败"))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func BulkCreateCategory(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		api_helper.HandleError(c, err)
		return
	}
	//read file
	req, err := utils.ReadCsv(fileHeader)
	if err != nil {
		api_helper.HandleError(c, err)
		return
	}
	// grpc 调用
	resp, err := rpc.BulkCreateCategory(c, req)
	if err != nil {
		api_helper.HandleError(c, errors.New("category rpc 调用失败"))
	}
	c.JSON(http.StatusOK, resp.BulkResponse)
}

func GetCategories(c *gin.Context) {
	page := utils.NewFromGinRequest(c, -1)
	// grpc调用
	resp, err := rpc.GetCategories(c, page)
	if err != nil {
		api_helper.HandleError(c, errors.New("category rpc 调用失败"))
	}
	c.JSON(http.StatusOK, resp.BulkResponse)
}
