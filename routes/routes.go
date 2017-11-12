package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/drcalx/record-holder-go/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/albums", handlers.List)
	router.GET("/albums/new", handlers.New)
}