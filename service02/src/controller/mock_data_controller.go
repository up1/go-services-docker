package controller

import (
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
)

type MockDataController struct {
	router *gin.Engine
}

func NewMockDataController(router *gin.Engine) MockDataController {
	return MockDataController{router}
}

func (mdc *MockDataController) CreateRouting() {
	// TODO add middleware
	// Define endpoint
	v1 := mdc.router.Group("/api/v1/mocks")
	{
		v1.GET("/", mdc.MockData)
	}
}

func (mdc *MockDataController) MockData(c *gin.Context) {
	items := []model.Item{
		{
			Id:   1,
			Name: "First item",
		},
		{
			Id:   2,
			Name: "Second item",
		},
	}
	c.JSON(http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"items": items})
}
