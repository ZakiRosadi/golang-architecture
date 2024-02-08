package migration

import (
	"fmt"
	"github-zaki-learning-golang/models"

	"gorm.io/gorm"
)

func DatabaseMigration(db *gorm.DB) {
	db.AutoMigrate(&models.BookModels{})

	fmt.Println("database migrated")
}