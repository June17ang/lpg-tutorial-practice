package controllers

import (
	"net/http"
	"tutorial-practice/02_mock_and_test/testable_code/services"

	"github.com/gin-gonic/gin"
)

var (
	PingController = pingController{}
)

type pingController struct{}

func (controller pingController) Ping(c *gin.Context) {
	result, err := services.PingService.HandlePing()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(http.StatusOK, result)
}
