package main

import (
	"controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	mdc := controller.NewMockDataController(router)
	mdc.CreateRouting()
	router.Run()
}
