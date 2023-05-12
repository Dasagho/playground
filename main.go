package main

import (
	"github.com/dasagho/playground/api/handler"
	"github.com/dasagho/playground/api/routers"
)

func main() {
	router := routers.SetupRouter()
	router.POST("/login", handler.PostLogin)
	router.Run(":8000")
}
