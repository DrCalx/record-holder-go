package main

import (
	"github.com/gin-gonic/gin"
	"github.com/DrCalx/record-holder-go/db"
	"github.com/DrCalx/record-holder-go/middlewares"
	"github.com/DrCalx/record-holder-go/routes"
	"github.com/DrCalx/record-holder-go/render"
)

var router *gin.Engine

func init() {
	db.Connect()
}

func main() {
	router = gin.Default()

	router.Use(middlewares.Connect)

	htmlRender := GinHTMLRender.New()
	router.HTMLRender = htmlRender.Create()

	routes.InitializeRoutes(router)

	router.Run()
}