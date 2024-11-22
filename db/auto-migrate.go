package db

import (
	"gorm.io/gorm"
	"ticketbookingapp/models"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{}, &models.Ticket{},&models.User{})
}