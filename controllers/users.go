package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/devardha/clivy-server/services"
	"github.com/gin-gonic/gin"
)

// GetUsers - get all users
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client := services.InitDB()
		collection := client.Database("clivy").Collection("users")

		fmt.Println(collection)

		c.JSON(200, gin.H{
			"users": "",
		})
	}
}

// RegisterUser - register new user
func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "true",
		})
	}
}
