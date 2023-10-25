package controllers

import (
	"github.com/fahimimam/UserStore/DB"
	"github.com/fahimimam/UserStore/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type SignUp struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

type TagsRegister struct {
	Names  []string `json:"names"`
	Expiry int64    `json:"expiry"`
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
	if err := DB.Get().Preload("Tags").Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Firstname + " " + user.Lastname,
		"phone": user.Phone,
		"tags":  user.Tags,
	})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User

	DB.Get().Preload("Tags").Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func AssignTags(c *gin.Context) {
	var body TagsRegister
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.Param("id")
	var user models.User
	DB.Get().Preload("Tags").First(&user, userID)

	duration := time.Duration(body.Expiry) * time.Millisecond
	for _, name := range body.Names {
		tag := models.Tags{
			Name:   name,
			Expiry: time.Now().Add(duration),
		}
		DB.Get().Create(&tag)
		user.Tags = append(user.Tags, tag)
	}
	DB.Get().Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "Tags added to the user",
	})
}

func SearchUsersByTags(c *gin.Context) {
	tags := c.Query("tags")
	tagList := strings.Split(tags, ",")
	userMap := make(map[uint]models.User)
	var users []models.User
	DB.Get().Where("tags.name IN (?)", tagList).
		Joins("JOIN user_tags ON user_tags.user_id = users.id").
		Joins("JOIN tags ON tags.id = user_tags.tags_id").
		Preload("Tags").
		Find(&users)

	for _, user := range users {
		userMap[user.ID] = user
	}

	c.JSON(http.StatusOK, gin.H{
		"users": userMap,
	})
}

func GetAllTags(c *gin.Context) {
	var tags []models.Tags

	DB.Get().Find(&tags)
	c.JSON(http.StatusOK, gin.H{
		"users": tags,
	})
}
