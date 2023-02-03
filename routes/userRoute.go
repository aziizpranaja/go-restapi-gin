package routes

import (
	"go-restapi-gin/controllers"
	"go-restapi-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(incomingRoutes *gin.Engine){
	user := incomingRoutes.Group("/user")
	user.POST("/", controllers.Register)
	user.POST("/login", controllers.Login)
	user.GET("/", middlewares.CheckCredetial, controllers.Profile)
	user.PUT("/", middlewares.CheckCredetial, controllers.ChangePass)
}