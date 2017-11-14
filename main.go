package main

import (
	"github.com/gin-gonic/gin"
	"github.com/DrCalx/record-holder-go/db"
	"github.com/DrCalx/record-holder-go/middlewares"
	"github.com/DrCalx/record-holder-go/routes"
)

var router *gin.Engine

func init() {
	db.Connect()
}

func main() {
	router = gin.Default()

	router.Use(middlewares.Connect)

	// router.LoadHTMLGlob("templates/*/*.html")

	htmlRender := GinHTMLRender.New()
	htmlRender.Debug = gin.IsDebugging()
	htmlRender.Layout = "layouts/default"

	router.HTMLRender = htmlRender.Create()

	routes.InitializeRoutes(router)

	router.Run()
}