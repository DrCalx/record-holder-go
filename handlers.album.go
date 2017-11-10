package main

import (
	"net/http"
	"strconv"

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

func getAlbum(c *gin.Context) {
	if albumID, err := strconv.Atoi(c.Param("album_id")); err == nil {
		if album, err := getAlbumByID(albumID); err == nil {
			c.HTML(
				http.StatusOK,
				"album.html",
				gin.H{
					"title": album.Title,
					"payload": album,
				},
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}