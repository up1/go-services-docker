package controller

import (
	"fmt"
	"gateway"
	"github.com/alexcesaro/statsd"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

type HelloController struct {
	s02Gateway gateway.Service02Gateway
	statsd     *statsd.Client
}

func NewHelloController(serviceName string, statsd *statsd.Client) HelloController {
	base, _ := url.Parse(fmt.Sprintf("http://%v", serviceName))
	return HelloController{
		s02Gateway: gateway.NewHttpServiceGateway(base),
		statsd:     statsd,
	}
}

func (hc *HelloController) Hello(c *gin.Context) {
	timing := hc.statsd.NewTiming()

	model, err := hc.s02Gateway.Inquiry()

	if err != nil {
		hc.statsd.Increment("helloFailure")
		c.JSON(http.StatusNotFound,
			gin.H{
				"code":    http.StatusNotFound,
				"message": err.Error()})
	} else {
		hc.statsd.Increment("helloSuccess")
		c.JSON(http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": model})
	}
	timing.Send("helloTiming")
}
