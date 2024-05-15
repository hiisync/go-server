package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hiisync/werix-server/internal/handlers"
)

// Routes sets up the routes for the server
func Routes(route *gin.Engine) {
	route.GET("/", handlers.HelloWorld)
}
