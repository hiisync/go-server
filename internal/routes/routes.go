package routes

import (
	"lunar-server/internal/handlers"

	"github.com/gin-gonic/gin"
)

// Routes sets up the routes for the server
func Routes(route *gin.Engine) {
	api := route.Group("/api")

	users := api.Group("/users")
	users.POST("/create", handlers.CreateUser)

}
