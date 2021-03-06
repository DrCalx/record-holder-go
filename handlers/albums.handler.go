package handlers

import (
	"net/http"
	// "strconv"

	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
	"github.com/DrCalx/record-holder-go/models"
)

func New(c *gin.Context) {
	album := models.Album{}

	c.HTML(
		http.StatusOK,
		"albums/form",
		gin.H{
			"album": album,
		},
	)
}

// func showIndexPage(c *gin.Context) {
// 	albums := getAllAlbums()

// 	c.HTML(
// 		http.StatusOK,
// 		"index.html",
// 		gin.H{
// 			"title": "Record Holder",
// 			"albums": albums,
// 		},
// 	)
// }

// func getAlbum(c *gin.Context) {
// 	if albumID, err := strconv.Atoi(c.Param("album_id")); err == nil {
// 		if album, err := getAlbumByID(albumID); err == nil {
// 			c.HTML(
// 				http.StatusOK,
// 				"album.html",
// 				gin.H{
// 					"title": album.Title,
// 					"album": album,
// 				},
// 			)
// 		} else {
// 			c.AbortWithError(http.StatusNotFound, err)
// 		}
// 	} else {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}
// }

func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	albums := []models.Album{}
	err := db.C(models.CollectionAlbum).Find(nil).All(&albums)
	if err != nil {
		c.Error(err)
	}
	c.HTML(
		http.StatusOK,
		"albums/list",
		gin.H{
			"title": "Record Holder - All Albums",
			"albums": albums,
		},
	)
}

func Edit(c *gin.Context) {

}

func Delete(c *gin.Context) {
	
}

func Save(c *gin.Context) {

}

func View(c *gin.Context) {
	
}