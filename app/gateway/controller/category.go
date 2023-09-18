package controller

import (
	"errors"
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
		api_helper.HandleError(c, errors.New("结构体绑定失败"))
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

func GetCategories(c *gin.Context) {
	page := utils.NewFromGinRequest(c, -1)
	categoryPage := utils.NewCategoryPages(page)
	// grpc调用
	resp, err := rpc.GetCategories(c, categoryPage)
	if err != nil {
		api_helper.HandleError(c, errors.New("category rpc 调用失败"))
	}
	c.JSON(http.StatusOK, resp.BulkResponse)
}
