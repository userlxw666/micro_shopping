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
		categoryGroup.POST("/create", controller.CreateCategory)
		categoryGroup.GET("/get", controller.GetCategories)
	}
	productGroup := ginRouter.Group("/product")
	productGroup.Use(middleware.MiddleJWT)
	{
		productGroup.POST("/create", controller.CreateProduct)
		productGroup.GET("/get", controller.GetProducts)
		productGroup.PUT("/update", controller.UpdateProduct)
		productGroup.DELETE("/delete", controller.DeleteProduct)
	}
	cartGroup := ginRouter.Group("/cart")
	cartGroup.Use(middleware.MiddleJWT)
	{
		cartGroup.GET("/get", controller.GetCart)
		cartGroup.POST("/add", controller.AddItem)
		cartGroup.PUT("/update", controller.UpdateItem)
	}
	orderGroup := ginRouter.Group("/order")
	orderGroup.Use(middleware.MiddleJWT)
	orderGroup.POST("/complete", controller.CompleteOrder)
	orderGroup.POST("/cancel", controller.CancelOrder)
	orderGroup.GET("/get", controller.GetOrders)
	ginRouter.Run()
}
