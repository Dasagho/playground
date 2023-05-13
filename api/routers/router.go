package routers

import (
	handler "github.com/dasagho/playground/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", handler.GetUsersLoged)
	router.GET("/users/:id", handler.GetUser)
	router.GET("subjects/:id", handler.GetResources)
	router.POST("/login", handler.PostLogin)
	return router
}
