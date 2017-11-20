package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/DrCalx/record-holder-go/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/", handlers.Home)

	router.GET("/albums", handlers.List)
	router.GET("/albums/new", handlers.New)
	router.POST("/albums/save", handlers.Save)
	router.GET("/albums/view/:id", handlers.View)
	router.GET("/albums/edit", handlers.Edit)
	router.POST("/albums/delete", handlers.Delete)
}