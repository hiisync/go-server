package server

import (
	"lunar-server/internal/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	routes.Routes(r)

	host := os.Getenv("SERVER_ADDRESS")
	r.Run(host)
}
