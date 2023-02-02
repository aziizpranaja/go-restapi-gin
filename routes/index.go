package routes

import (
	"github.com/gin-gonic/gin"
)

func AllRoutes() {
	route := gin.New()
	route.Use(gin.Logger())

	// v1 := route.Group("/v1")
	UserRoute(route)
	ProductRoute(route)

	route.Run()
}