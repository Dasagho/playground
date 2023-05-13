package routers

import (
	"github.com/dasagho/playground/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", handler.GetUsersLoged)
	router.POST("/login", handler.PostLogin)
	return router
}
