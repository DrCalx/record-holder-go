package main

type album struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
}

var albumList = []album{
	album{ID: 1, Title: "Please Please Me", Artist: "The Beatles"},
	album{ID: 2, Title: "With the Beatles", Artist: "The Beatles"},
}

func getAllAlbums() []album {
	return albumList
}