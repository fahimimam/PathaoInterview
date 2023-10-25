package main

import (
	"github.com/fahimimam/UserStore/DB"
	"github.com/fahimimam/UserStore/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/users/register", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUser)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/tags", controllers.GetAllTags)

	router.POST("/users/:id/tags", controllers.AssignTags)
	router.GET("/users/tags", controllers.SearchUsersByTags)

	router.GET("/migrate", DB.AutoMigrateDB)
	router.Run("localhost:3000")
}
