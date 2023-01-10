package main

import (
	"github.com/gin-gonic/gin"
	"go-web/interactuando-con-la-api/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/ping", handlers.Ping)

	products := router.Group("/products")

	products.GET("/", handlers.GetAllProducts)
	products.GET("/:id", handlers.GetProductById)
	products.GET("/search", handlers.GetProductsByPrice)

	products.POST("/", handlers.AddProduct)

	router.Run()
}
