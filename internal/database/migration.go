package database

import (
	"lunar-server/internal/database/models"
)

func Migrate() {
	DB.AutoMigrate(&models.User{})
}
