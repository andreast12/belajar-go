package controllers

import (
	"example/belajar-go/models"
	"example/belajar-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Name string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	utils.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"message": "Sign up successful"})
}

func SignIn(c *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var user models.User
	result := utils.DB.Where("email = ?", body.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect email or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect email or password"})
		return
	}

	token := utils.GenerateToken(user.ID)	
	c.SetCookie("jwt", token, 3600 * 24 * 30, "", "", utils.IsInProduction(), true)
	c.JSON(http.StatusOK, gin.H{"message": "Sign in successful"})
}