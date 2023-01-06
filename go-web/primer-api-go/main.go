package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	type FullName struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
	//Ejercicio 1
	router.GET("/ping", func(context *gin.Context) {
		context.String(200, "Pong")
	})
	//Ejercicio 2
	router.POST("/Saludar", func(context *gin.Context) {
		body := FullName{}
		err := context.ShouldBind(&body)
		if err != nil {
			context.AbortWithStatus(400)
			return
		}
		context.String(200, "Hola"+" "+body.FirstName+" "+body.LastName)
	})

	router.Run()

}
