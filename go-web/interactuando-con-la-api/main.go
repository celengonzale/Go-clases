package main

import (
	"github.com/gin-gonic/gin"
	"go-web/interactuando-con-la-api/handlers"
)

func main() {
	router := gin.Default()
	products := router.Group("/products")

	router.GET("/ping", handlers.Ping)
	products.GET("/", handlers.GetAllProducts)
	products.GET("/:id", handlers.GetProductById)
	products.GET("/search", handlers.GetProductsByPrice)
	router.Run()
}
