package routers

import (
	"github.com/devardha/clivy-server/controllers"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes - grouping all apis
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("v1")
	{
		api.GET("/users", controllers.GetUsers())
	}
}
