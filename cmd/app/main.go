package main

import (
	"lunar-server/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app.Run()
}
