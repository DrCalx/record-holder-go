package main

import (
	"errors"
)

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

func getAlbumByID(id int) (*album, error) {
	for _, a := range albumList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Album not found")
}