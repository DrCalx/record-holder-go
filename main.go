package main

import (
	"github.com/gin-gonic/gin"
	"github.com/drcalx/record-holder-go/db"
	"github.com/drcalx/record-holder-go/middlewares"
	"github.com/drcalx/record-holder-go/routes"
)

var router *gin.Engine

func init() {
	db.Connect()
}

func main() {
	router = gin.Default()

	router.Use(middlewares.Connect)

	router.LoadHTMLGlob("templates/*")

	routes.InitializeRoutes(router)

	router.Run()
}