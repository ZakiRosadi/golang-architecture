package main

import (
	"github-zaki-learning-golang/controllers"
	"github-zaki-learning-golang/database"
	"github-zaki-learning-golang/routes"

	// "github-zaki-learning-golang/models"
	"github-zaki-learning-golang/repository"

	"github.com/gin-gonic/gin"
)


func main(){
	
	
	dataDB :=database.DBConnection()
	//repository
	bookRepository := repository.NewBookRepository(dataDB)
	

	//controllers
	bookController := controllers.NewBookController(bookRepository)

	
	//router
	router := gin.Default()
	routes.NewBookRoutes(router, bookController)
	

	


	router.Run(":8080")
}






