package database

import (
	"gh22-go/internal/entities/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Token{})
	db.AutoMigrate(&model.User{})
}
