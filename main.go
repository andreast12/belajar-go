package main

import (
	"example/belajar-go/controllers"
	"example/belajar-go/middlewares"
	"example/belajar-go/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnvVars()
	utils.ConnectDB()
}

func main() {
	router := gin.Default()

	todos := router.Group("/todos")
	todos.Use(middlewares.RequiresAuth())
	{
		todos.GET("/", controllers.GetTodos)
		todos.GET("/:id", controllers.GetTodoById)
		todos.POST("/", controllers.CreateTodo)
		todos.PATCH("/:id", controllers.ToggleTodoCompleted)
		todos.DELETE("/:id", controllers.DeleteTodo)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/", controllers.SignIn)
	}

	router.Run()
}