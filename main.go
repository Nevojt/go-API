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
	app.GET("/user/:id", controllers.GetUserById)
	app.DELETE("/user/:id", controllers.DeleteUser)
	app.PATCH("/user/:id", controllers.UpdateUser)

	if err := app.Run(); err != nil {
		fmt.Println("Failed to start server:", err)
		os.Exit(1)
	}
}
