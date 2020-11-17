package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var(
	router = gin.Default()
)

func main()  {
	router.GET("/ping", pingHander);
	router.Run(":8080")
}

func pingHander(c *gin.Context){
	c.String(http.StatusOK, "pong")
}