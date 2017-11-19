package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"layout",
		gin.H{
			"title": "Record Holder",
		},
	)
}