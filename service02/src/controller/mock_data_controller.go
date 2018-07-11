package controller

import (
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
)

type MockDataController struct {
}

func NewMockDataController() MockDataController {
	return MockDataController{}
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
