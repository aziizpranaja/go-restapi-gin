package main

import (
	"go-restapi-gin/controllers"
	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	models.ConnectionDatabase()

	route.GET("/", controllers.Index)
	route.GET("/:id", controllers.Show)
	route.POST("/", controllers.Create)
	route.PUT("/:id", controllers.Update)
	route.DELETE("/", controllers.Delete)

	route.Run()
}