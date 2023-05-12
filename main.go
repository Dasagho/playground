package main

import (
	"github.com/dasagho/playground/api/routers"
)

func main() {
	router := routers.SetupRouter()
	router.Run(":8000")
}
