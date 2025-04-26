package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// Public auth routes
	auth := api.Group("/auth")
	{
		auth.POST("/register", Register)
		auth.POST("/login", Login)
	}

	// Protected order routes
	protected := api.Group("/orders")
	protected.Use(AuthMiddleware())
	{
		protected.POST("/", CreateOrder)
		protected.GET("/", GetAllOrders)
		protected.GET("/:id", GetOrderByID)
		protected.PUT("/:id", UpdateOrder)
		protected.DELETE("/:id", DeleteOrder)
	}
}
