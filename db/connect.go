package db

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
	Mongo *mgo.DialInfo
)

func Connect() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		fmt.Println("No connection string provided. Set MONGODB_URI")
	}
	uri = strings.TrimSuffix(uri, "?ssl=true")

	dialInfo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Println("database connection failed:", err)
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("database connection successful")
	Session = s
	Mongo = dialInfo
}