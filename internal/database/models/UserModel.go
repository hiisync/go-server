package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"not null;size:255" json:"name"`
	Username string `gorm:"unique;not null;size:255" json:"username"`
	Email    string `gorm:"unique;not null;size:255" json:"email"`
	Password string `gorm:"not null;size:255" json:"password"`
}
