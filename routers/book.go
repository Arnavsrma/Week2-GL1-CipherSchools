package router

import (
	"github.com/Arnavsrma/Week2-GL1-Cipherschool/database"
	"github.com/Arnavsrma/Week2-GL1-Cipherschool/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default() //get the defult engine for further customization
	api := handler.Handler{
		DB: database.GetDB(), //set the handler DB
	}
	router.GET("/books", api.GetBooks) //set the function for http://localhost:8080/books : Get request
	//while calling handler function, gin will pass the *gin.Context in the handler function

	router.POST("/books", api.SaveBook)

	return router
}
