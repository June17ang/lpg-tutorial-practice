package main

import (
	"github.com/gin-gonic/gin"
	"tutorial-practice/02_mock_and_test/testable_code/controllers"
)

var(
	router = gin.Default()
)

func main()  {
	router.GET("/ping", controllers.Ping)

	router.Run(":8080")
}