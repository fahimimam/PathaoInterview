package main

import (
	"github.com/fahimimam/UserStore/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/users/register", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUser)

	router.Run("localhost:3000")
}
