package server

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	Routes(r)

	r.Run()
}
