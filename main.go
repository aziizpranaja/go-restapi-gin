package main

import (
	"go-restapi-gin/controllers"
	"go-restapi-gin/initializers"
	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVar()
	models.ConnectionDatabase()
}

func main() {
	route := gin.Default()
	
	// v1 := route.Group("/v1")
	// Route Product
	product := route.Group("/product")
	product.GET("/", controllers.Index)
	product.GET("/:id", controllers.Show)
	product.POST("/", controllers.Create)
	product.PUT("/:id", controllers.Update)
	product.DELETE("/", controllers.Delete)

	// Route User
	user := route.Group("/user")
	user.POST("/", controllers.Register)
	user.POST("/login", controllers.Login)

	route.Run()
}