package routes

import (
	"lunar-server/internal/handlers"
	"lunar-server/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Routes sets up the routes for the server
func Routes(route *gin.Engine) {
	api := route.Group("/api")

	users := api.Group("/users")
	users.POST("/create", handlers.CreateUser)
	users.POST("/login", handlers.UserLogin)
	users.GET("/validate", middleware.RequireAuth, handlers.ValidateUser)
}
