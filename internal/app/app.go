package app

import (
	"lunar-server/internal/database"
	"lunar-server/internal/server"
)

// Run starts the server
func Run() {
	database.InitDB()
	database.Migrate()
	server.StartServer()
}
