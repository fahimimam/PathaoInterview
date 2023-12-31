package controllers

import (
	"github.com/fahimimam/UserStore/DB"
	"github.com/fahimimam/UserStore/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func UserCreation(c *gin.Context, user *models.User) {
	pass, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal("failed to hash Password ", err)
	}
	user.Password = pass
	DB.Get().Save(user)
}

func UserUpdate(c *gin.Context, user *models.User, fields map[string]interface{}) {
	pass, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal("failed to hash Password ", err)
	}
	user.Password = pass
	DB.Get().Model(&user).Updates(fields)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func isTagExpired(expirationTime time.Time) bool {
	return time.Now().After(expirationTime)
}
func deleteExpiredTags() {
	var expiredTags []models.Tags
	DB.Get().Where("expiry < ?", time.Now()).Find(&expiredTags)
	for _, tag := range expiredTags {
		DB.Get().Delete(&tag)
	}
	var users []models.User
	DB.Get().Find(&users)
	for _, user := range users {
		if user.Tags != nil {
			DB.Get().Model(user).Association("Tags").Replace(user.Tags)
		}
	}
	return
}
