package main

import (
	"github.com/gin-gonic/gin"
	"github.com/DrCalx/record-holder-go/db"
	"github.com/DrCalx/record-holder-go/middlewares"
	"github.com/DrCalx/record-holder-go/routes"
	// eztemplate "github.com/michelloworld/ez-gin-template"
)

var router *gin.Engine

func init() {
	db.Connect()
}

func main() {
	router = gin.Default()

	router.Use(middlewares.Connect)

	// render := eztemplate.New()
	// render.TemplatesDir = "templates/"
	// render.Layout = "layouts/base"
	// render.Debug = true
	// router.HTMLRender = render.Init()

	router.LoadHTMLGlob("views/**/*.html")

	routes.InitializeRoutes(router)

	router.Run()
}