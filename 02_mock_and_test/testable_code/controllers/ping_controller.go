package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tutorial-practice/02_mock_and_test/testable_code/services"
)

func Ping(c *gin.Context){
	c.String(http.StatusOK, services.HandlePing())
}