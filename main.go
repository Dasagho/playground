package main

import (
	"fmt"
	"os"

	"github.com/dasagho/playground/api/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Failed loading env file")
	}
}

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))

	router := routers.SetupRouter()
	router.Run(":8000")
}
