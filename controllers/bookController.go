package controllers

import (
	"fmt"
	"github-zaki-learning-golang/dto"
	"github-zaki-learning-golang/models"
	"github-zaki-learning-golang/repository"

	"github.com/gin-gonic/gin"
)

type IBookController interface {
	FindAll() []models.BookModels
	FindOneById(id int) (models.BookModels, error)
	Create(req dto.BookRequest) (dto.BookResponse, error)
	Update(id int, req dto.BookUpdateRequest) (dto.BookUpdateResponse, error)
	Delete(id int, req dto.BookDeleteRequest) (dto.BookDeleteResponse, error)
}

type bookController struct {
	bookrepository repository.IBookRepository
}

func NewBookController(bookrepository repository.IBookRepository) IBookController{
	return &bookController{
		bookrepository :bookrepository,
	}
}

// func (r *bookController) Findall() []models.BookModels{
// 	err := r.bookrepository.FindAll()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return err
// }

//shortcut
func (con *bookController) FindAll() []models.BookModels{
	return con.bookrepository.FindAll()
}

func (con *bookController) FindOneById(id int) (models.BookModels, error){
	data,err := con.bookrepository.FindOneById(id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (con *bookController) Create(req dto.BookRequest) (dto.BookResponse, error){
	// var request models.BookModels
	request := models.BookModels{
		Name: req.Name,
		Price: req.Price,
		Author: req.Author,
		Description: req.Description,
	}
	data, err := con.bookrepository.Create(request)
	if err != nil {
		fmt.Println("============= cannot create ============", err)
	}
	return dto.BookResponse{
		Name: data.Name,
		Price: data.Price,
		Author: data.Author,
		Description: data.Description,
	}, nil
}

func (con *bookController) Update(id int, req dto.BookUpdateRequest) (dto.BookUpdateResponse, error){
	
	idData, err := con.bookrepository.FindOneById(id)
	if err != nil {
		fmt.Println("============= cannot update ============", err)
	}
	//idData.Name = req.Name means which data needs to be updated with new request
	idData.Name = req.Name
	idData.Price = req.Price
	idData.Author = req.Author
	idData.Description = req.Description

	con.bookrepository.Update(idData)
	return dto.BookUpdateResponse{
		Name: idData.Name,
		Price: idData.Price,
		Author: idData.Author,
		Description: idData.Description,
	}, nil
	
}

func (con *bookController) Delete(id int, req dto.BookDeleteRequest) (dto.BookDeleteResponse, error){
	data, _ := con.bookrepository.FindOneById(id)

	book, err := con.bookrepository.Delete(data)
	if err != nil {
		fmt.Println("============= cannot delete ============", err)
	}
	return dto.BookDeleteResponse{
		Name: book.Name,
		Price: book.Price,
		Author: book.Author,
		Description: book.Description,
	}, nil
}

















func TestController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func GetIdController(c *gin.Context) {
	id := c.Param("id")
	var name = c.Param("name")
	c.JSON(200, gin.H{
		"id":   id,
		"name": name,
	})
}

func GetIdControllerQuery(c *gin.Context) {
	title := c.Query("title")
	lastname := c.Query("lname")
	c.JSON(200, gin.H{
		"judul":        title,
		"namaterakhir": lastname,
	})
}

func PostController (c *gin.Context) {
	//need to make struct first if we want to post data.

	var reqData models.BookInput //data is still zero, no value.

	err := c.ShouldBindJSON(&reqData)//data will be filled in form.
	if err != nil {
		// fmt.Println("error message : this is error", "error :", err)
		c.JSON(500, gin.H{
			"error message": "this is error",
			"err" : err.Error(),
		})
		return 
		}


	
	c.JSON(200, gin.H{
		"name": reqData.Name,
		"price": reqData.Price,
		"author_kita" : reqData.AuthorKita,

	})

}