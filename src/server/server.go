package server

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago18l/restful-gin-api/src/models"
	"github.com/thiago18l/restful-gin-api/src/routes"
)

func Start() {
	db := models.SetupModels()
	router := gin.New()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	routes.Routes(router)
	run(router)
}

func run(router *gin.Engine) {
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
