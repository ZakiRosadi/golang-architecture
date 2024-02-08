package repository

import (
	"errors"
	"fmt"
	"github-zaki-learning-golang/models"

	"gorm.io/gorm"
)

type IBookRepository interface {
	FindAll() []models.BookModels
	FindOneById(id int) (models.BookModels, error)
	Create(book models.BookModels) (models.BookModels, error)
	Update(book models.BookModels) (models.BookModels, error)
	Delete(book models.BookModels) (models.BookModels, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository{
	return &bookRepository{
		db :db,}

}

func (r *bookRepository) FindAll() []models.BookModels {
	var books []models.BookModels
	err := r.db.Find(&books).Error
	if err != nil {
		panic(err)
	}
	return books
}


func (r *bookRepository) FindOneById(id int) (models.BookModels, error) {
	var books models.BookModels
	err := r.db.Where("id = ?", id).First(&books).Error
	fmt.Println("errorrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr : ", err)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return books, errors.New("book not found")
		}
		return books, err
	}


	return books, nil
}

func (r *bookRepository) Create(book models.BookModels) (models.BookModels, error){
	err := r.db.Create(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *bookRepository) Update(book models.BookModels) (models.BookModels, error){
	err := r.db.Save(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *bookRepository) Delete(book models.BookModels) (models.BookModels, error){
	err := r.db.Delete(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}