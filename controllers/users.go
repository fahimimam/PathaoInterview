package controllers

import (
	"github.com/fahimimam/UserStore/DB"
	"github.com/fahimimam/UserStore/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUp struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

func CreateUser(c *gin.Context) {
	var body SignUp

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := &models.User{
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Password:  body.Password,
		Phone:     body.Phone,
	}
	go UserCreation(c, user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully created",
		"name":    user.Firstname + " " + user.Lastname,
	})
}

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := DB.Get().Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Firstname + " " + user.Lastname,
		"phone": user.Phone,
	})
}
