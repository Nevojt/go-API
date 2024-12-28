package main

import (
	"api/pkg/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	app := gin.New()
	app.POST("/user", controllers.CreateUser)
	app.GET("/user", controllers.GetAllUsers)

	if err := app.Run(); err != nil {
		fmt.Println("Failed to start server:", err)
		os.Exit(1)
	}
}
