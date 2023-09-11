package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"micro_shopping/app/gateway/rpc"
	"micro_shopping/idl/pb"
	"micro_shopping/pkg/api_helper"
	"net/http"
)

func TestPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "pong",
	})
}

func UserRegister(c *gin.Context) {
	var req pb.UserRequest
	if err := c.ShouldBind(&req); err != nil {
		api_helper.ResponseHandler(c, err)
		return
	}
	userResp, err := rpc.UserRegister(c, &req)
	if err != nil {
		api_helper.HandleError(c, errors.New("grpc调用失败"))
		return
	}
	c.JSON(200, userResp.UserDetail)
}

func UserLogin(c *gin.Context) {
	var req pb.UserRequest
	if err := c.ShouldBind(&req); err != nil {
		api_helper.HandleError(c, err)
		return
	}

	userResp, err := rpc.UserLogin(c, &req)
	if err != nil {
		api_helper.HandleError(c, errors.New("grpc调用失败"))
		return
	}
	api_helper.ResponseHandler(c, userResp)
}
