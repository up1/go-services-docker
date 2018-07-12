package main

import (
	"controller"
	"github.com/alexcesaro/statsd"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	hc := controller.NewHelloController(os.Getenv("SERVICE02"))
	router := gin.Default()
	v1 := router.Group("/api/v1/hello")
	{
		v1.GET("/", hc.Hello)
	}
	router.Run()
}

func createStatsDClient(address string) (*statsd.Client, error) {
	return statsd.New(statsd.Address(address))
}
