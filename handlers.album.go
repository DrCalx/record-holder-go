package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	albums := getAllAlbums()

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Record Holder",
			"payload": albums,
		},
	)
}