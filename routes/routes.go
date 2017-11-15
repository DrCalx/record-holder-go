package routes

import (
	"github.com/gin-gonic/gin"
	// "github.com/DrCalx/record-holder-go/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			200,
			"layout",
			gin.H{
				"title": "hello world",
				"text": "hi there",
			},
		)
	})
	// router.GET("/albums", handlers.List)
	// router.GET("/albums/new", handlers.New)
}