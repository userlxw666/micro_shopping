package router

import (
	"github.com/gin-gonic/gin"
	"micro_shopping/app/gateway/controller"
	"micro_shopping/app/gateway/middleware"
)

func NewRouter() {
	ginRouter := gin.Default()

	userGroup := ginRouter.Group("/user")
	{
		userGroup.POST("/register", controller.UserRegister)
		userGroup.POST("/login", controller.UserLogin)
	}
	categoryGroup := ginRouter.Group("/category")
	categoryGroup.Use(middleware.MiddleJWT)
	{
		categoryGroup.GET("/ping", controller.TestPing)
		categoryGroup.POST("/create", controller.CreateCategory)
		categoryGroup.GET("/get", controller.GetCategories)
		categoryGroup.GET("/bulkcreate", controller.BulkCreateCategory)
	}
	ginRouter.Run()
}
