package controllers

import (
	"example/belajar-go/models"
	"example/belajar-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	user, _ := c.Get("user")
	utils.DB.Where("user_id = ?", user.(models.User).ID).Find(&todos)
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func GetTodoById(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	user, _ := c.Get("user")
	result := utils.DB.Where("user_id = ?", user.(models.User).ID).First(&todo, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func CreateTodo(c *gin.Context) {
	var body struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, _ := c.Get("user")
	todo := models.Todo{Name: body.Name, Completed: false, UserID: user.(models.User).ID}
	utils.DB.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{"todo": todo})
}

func ToggleTodoCompleted(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	result := utils.DB.First(&todo, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": result.Error.Error()})
		return
	}

	user, _ := c.Get("user")
	if todo.UserID != user.(models.User).ID {
		c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
		return
	}

	utils.DB.Model(&todo).Update("completed", !todo.Completed)
	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	result := utils.DB.First(&todo, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": result.Error.Error()})
		return
	}

	user, _ := c.Get("user")
	if todo.UserID != user.(models.User).ID {
		c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
		return
	}

	utils.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"todo": todo})
}