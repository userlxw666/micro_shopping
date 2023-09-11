package router

import (
	"github.com/gin-gonic/gin"
	"micro_shopping/app/gateway/controller"
)

func NewRouter() {
	ginRouter := gin.Default()
	ginRouter.GET("/ping", controller.TestPing)
	ginRouter.POST("/register", controller.UserRegister)
	ginRouter.POST("/login", controller.UserLogin)
	ginRouter.Run()
}
