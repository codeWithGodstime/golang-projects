package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	protected := api.Group("/orders")
	protected.Use(AuthMiddleware())

	{
		protected.POST("/", CreateOrder)
		protected.GET("/", GetAllOrders)
		protected.GET("/:id", GetOrderByID)
		protected.PUT("/:id", UpdateOrder)
		protected.DELETE("/:id", DeleteOrder)
	}

	{
		api.POST("/auth/register", Register)
		api.POST("/auth/login", Login)
	}
}
