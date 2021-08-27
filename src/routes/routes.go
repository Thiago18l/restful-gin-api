package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/thiago18l/restful-gin-api/src/controllers"
)

// Routes define the routes of the application
func Routes(router *gin.Engine) {
	router.GET("/books", controllers.FindBook)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.FindBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
}
