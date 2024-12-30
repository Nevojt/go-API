package main

import (
	"api/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port := os.Getenv("RUN_APP_PORT")
	//gin.SetMode(os.Getenv("GIN_MODE"))

	app := gin.New()
	app.POST("/user", controllers.CreateUser)
	app.GET("/user/:id", controllers.GetUserById)
	app.POST("/login", controllers.LoginHandler)

	protected := app.Group("/protected")
	protected.Use(controllers.AuthMiddleware())
	{
		protected.GET("/users", controllers.GetAllUsers)
		protected.PATCH("/user/:id", controllers.UpdateUser)
		protected.DELETE("/user/:id", controllers.DeleteUser)
	}

	if err := app.Run(port); err != nil {
		fmt.Println("Failed to start server:", err)
		os.Exit(1)
	}
	fmt.Println("Server started on port", port)
}
