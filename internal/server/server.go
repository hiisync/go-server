package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hiisync/werix-server/internal/routes"
)

func StartServer() {
	r := gin.Default()

	routes.Routes(r)

	r.Run()
}
