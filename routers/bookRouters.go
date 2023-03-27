package routers

import (
	"chapter2_1/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)
	router.PUT("/book/:bookID", controllers.UpdateBook)
	router.GET("/book/:bookID", controllers.GetBook)
	router.DELETE("/book/:bookID", controllers.DeleteBook)

	return router
}
