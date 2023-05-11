package main

import (
	"github.com/dasagho/playground/api/handler"
)

func main() {
	router := handler.SetupRouter()
	router.Run(":8000")
}
