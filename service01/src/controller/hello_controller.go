package controller

import (
	"fmt"
	"gateway"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

type HelloController struct {
	s02Gateway gateway.Service02Gateway
}

func NewHelloController(serviceName string) HelloController {
	base, _ := url.Parse(fmt.Sprintf("http://%v", serviceName))
	return HelloController{
		s02Gateway: gateway.NewHttpServiceGateway(base),
	}
}

func (hc *HelloController) Hello(c *gin.Context) {
	model, err := hc.s02Gateway.Inquiry()

	if err != nil {
		c.JSON(http.StatusNotFound,
			gin.H{
				"code":    http.StatusNotFound,
				"message": err.Error()})
	} else {
		c.JSON(http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": model})
	}
}
