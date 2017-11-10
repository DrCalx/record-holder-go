package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

type Album struct {
	Id string
	Title string
	Artist string
}

// var templates = template.Must(template.ParseFiles("views/layout.html", "views/view.html", "views/edit.html"))

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	a := &Album{Id: "1", Title: "Please Please Me", Artist: "The Beatles"}
// 	renderTemplate(w, "view", a)
// }

// func editHandler(w http.ResponseWriter, r *http.Request) {
// 	a := &Album{Id: "1", Title: "Please Please Me", Artist: "The Beatles"}
// 	renderTemplate(w, "edit", a)
// }

// func renderTemplate(w http.ResponseWriter, tmpl string, a *Album) {
// 	t, _ := template.ParseFiles("views/layout.html", "views/"+tmpl+".html")
// 	err := t.ExecuteTemplate(w, "layout", a)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

func main() {
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	router.Run()
}