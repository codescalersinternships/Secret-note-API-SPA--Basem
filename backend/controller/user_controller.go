package controllers

import (
	"fmt"
	"net/http"

	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/database"
	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/models"
	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/utils"
	"github.com/gin-gonic/gin"
)

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserController struct {
	DB *database.DB
}

func NewUserController(db *database.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) Register(c *gin.Context) {
	var input AuthInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: input.Username, Password: input.Password}
	fmt.Print(user.Password)
	if err := uc.DB.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	fmt.Print(user.Username)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (uc *UserController) Login(c *gin.Context) {

	var input AuthInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.DB.GetUserByUsernameAndPassword(input.Username, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	tokenString, err := utils.GenerateToken(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}
