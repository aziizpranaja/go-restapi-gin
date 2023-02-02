package routes

import (
	"go-restapi-gin/controllers"
	"go-restapi-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoute(incomingRoutes *gin.Engine) {
	product := incomingRoutes.Group("/product")
	product.GET("/", middlewares.CheckCredetial, controllers.Index)
	product.GET("/:id", controllers.Show)
	product.POST("/", controllers.Create)
	product.PUT("/:id", controllers.Update)
	product.DELETE("/", controllers.Delete)
}