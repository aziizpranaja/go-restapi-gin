package main

import (
	"go-restapi-gin/initializers"
	"go-restapi-gin/models"
	"go-restapi-gin/routes"
)

func init(){
	initializers.LoadEnvVar()
	models.ConnectionDatabase()
}

func main() {
	routes.AllRoutes()
}