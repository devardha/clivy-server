package controllers

import "github.com/gin-gonic/gin"

// GetUsers - get all users
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"users": "users",
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
