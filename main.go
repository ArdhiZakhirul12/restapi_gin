package main

import (
	"rest_api_golang/models"
	"rest_api_golang/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	route.RegisterRoutes(r)

	r.Run()
}
