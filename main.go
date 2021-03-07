package main

import (
	"github.com/devardha/clivy-server/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	routers.ApplyRoutes(r)

	r.Run()
}
