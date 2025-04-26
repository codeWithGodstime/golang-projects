package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	{
		api.POST("/orders", CreateOrder)
		api.GET("/orders", GetAllOrders)
		api.GET("/orders/:id", GetOrderByID)
		api.PUT("/orders/:id", UpdateOrder)
		api.DELETE("/orders/:id", DeleteOrder)
	}
}
