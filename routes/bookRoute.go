package routes

import (
	"github-zaki-learning-golang/controllers"
	"github-zaki-learning-golang/dto"

	"strconv"

	"github.com/gin-gonic/gin"
)

type bookRoutes struct {
	bookController controllers.IBookController
}

func NewBookRoutes(gin *gin.Engine, bookController controllers.IBookController){
	r := bookRoutes{
		bookController : bookController,
	}
	group := gin.Group("/api/v1")
	group.GET("/books", r.FindAll)
	group.GET("/books/:id", r.FindOneById)
	group.POST("/books", r.Create)
	group.PATCH("/books/:id", r.Update)
	group.DELETE("/books/:id", r.Delete)
}





func (r *bookRoutes) FindAll(c *gin.Context){
	data:= r.bookController.FindAll()
	c.JSON(200, gin.H{
		"data" : data,
	})
}

func (r *bookRoutes) FindOneById(c *gin.Context){
	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams)
	
	if err != nil {
		c.JSON(400, gin.H{
			"message" : "invalid id",
			"error" : err,
		})
		return
	}

	data, err := r.bookController.FindOneById(id)
	
	if err != nil {
		c.JSON(404, gin.H{
			
			"message" : "cannot find id",
			"error" : err,
		})
		return
	}
	
	c.JSON(200, data)
}

func (r *bookRoutes) Create(c *gin.Context){
	var reqData dto.BookRequest

	err := c.ShouldBindJSON(&reqData)
	if err != nil {
		c.JSON(500, gin.H{
			"error" : err,
		})
		return
	}

	data , _ :=r.bookController.Create(reqData)
	 c.JSON(200, gin.H{
		"data" : data,
	})
}

func (r *bookRoutes) Update(c *gin.Context){
	idParams := c.Param("id")
	id, _ := strconv.Atoi(idParams)

	var reqData dto.BookUpdateRequest
	err := c.ShouldBindJSON(&reqData)
	if err != nil {
		c.JSON(500, gin.H{
			"error" : err,
		})
		return
	}

	data, _ := r.bookController.Update(id, reqData)
	c.JSON(200, gin.H{
		"data " : data,
	})
}

func (r *bookRoutes) Delete(c *gin.Context){
	idParams := c.Param("id")
	id, _ := strconv.Atoi(idParams)
	var reqData dto.BookDeleteRequest

	data,err := r.bookController.Delete(id, reqData)

	if err != nil {
		c.JSON(500, gin.H{
			"error" : err,
		})
		return
	}
	c.JSON(200, gin.H{
		"data " : data,
	})
}