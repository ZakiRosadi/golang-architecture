package dto

type BookRequest struct {
	Name        string `gorm:"not null" json:"name"`
	Price       int    `gorm:"type:int" json:"price"`
	Author      string `gorm:"type:varchar(50)" json:"author"`
	Description string `gorm:"type:text" json:"description"`
}

type BookResponse struct {
	Name        string `gorm:"not null" json:"name"`
	Price       int    `gorm:"type:int" json:"price"`
	Author      string `gorm:"type:varchar(50)" json:"author"`
	Description string `gorm:"type:text" json:"description"`
}

type BookUpdateRequest struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Price       int    `gorm:"type:int" json:"price"`
	Author      string `gorm:"type:varchar(50)" json:"author"`
	Description string `gorm:"type:text" json:"description"`
}

type BookUpdateResponse struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Price       int    `gorm:"type:int" json:"price"`
	Author      string `gorm:"type:varchar(50)" json:"author"`
	Description string `gorm:"type:text" json:"description"`
}

type BookDeleteRequest struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Price       int    `gorm:"type:int" json:"price"`
	Author      string `gorm:"type:varchar(50)" json:"author"`
	Description string `gorm:"type:text" json:"description"`
}

type BookDeleteResponse struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Price       int    `gorm:"type:int" json:"price"`
	Author      string `gorm:"type:varchar(50)" json:"author"`
	Description string `gorm:"type:text" json:"description"`
}