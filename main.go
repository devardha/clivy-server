package main

import (
	"context"
	"fmt"

	"github.com/devardha/clivy-server/routers"
	"github.com/devardha/clivy-server/services"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	client := services.InitDB()
	defer client.Disconnect(context.Background())

	routers.ApplyRoutes(r)

	r.GET("/test", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"test": "success",
		})
	})

	fmt.Println("database connected")
	r.Run()
}
