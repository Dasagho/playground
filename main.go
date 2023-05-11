package main

import (
	"github.com/dasagho/playground/api/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handler.LoginHandler)

	router.Run(":8000")
}
