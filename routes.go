package main

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/album/view/:album_id", getAlbum)
}