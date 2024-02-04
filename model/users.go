package model

import "gorm.io/gorm"

type User struct {
	gorm.Model        // Adds fields ID, CreatedAt, UpdatedAt, DeletedAt
	Email      string `gorm:"type:varchar(100);unique_index"`
	Name       string `gorm:"type:varchar(100)"`
	Password   string `gorm:"type:varchar(255)"` // Store hashed password
}
