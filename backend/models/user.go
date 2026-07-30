package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"default:'user'" json:"role"`
}
