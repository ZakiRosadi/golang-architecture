package models

import "time"

type BookModels struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Price       int    `gorm:"type:int" json:"price"`
	Author      string `gorm:"type:varchar(50)" json:"author"`
	Description string `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}