package main

import (
	"controller"
	"github.com/gin-gonic/gin"
)

func main() {
	mdc := controller.NewMockDataController()
	router := gin.Default()
	v1 := router.Group("/api/v1/mocks")
	{
		v1.GET("/", mdc.MockData)
	}
	router.Run()
}
