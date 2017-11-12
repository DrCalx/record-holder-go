package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionAlbum = "albums"
)

type Album struct {
	Id				bson.ObjectId	`json:"_id,omitempty" bson:"_id,omitempty"`
	Title			string				`json:"title" form:"title" binding:"required" bson:"title"`
	Artist		string				`json:"artist" form:"artist" binding:"required" bson:"artist"`
	CreatedOn	int64					`json:"created_on" bson:"created_on"`
	UpdatedOn	int64					`json:"updated_on" bson:"updated_on"`
}