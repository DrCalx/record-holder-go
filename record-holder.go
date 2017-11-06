package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Album struct {
	_id string
	Title string
	Artist string
}

var templates = template.Must(template.ParseFiles("views/view.html"))

func viewHandler(w http.ResponseWriter, r *http.Request) {
	a := &Album{_id: "1", Title: "Please Please Me", Artist: "The Beatles"}
	renderTemplate(w, "view", a)
}

func renderTemplate(w http.ResponseWriter, tmpl string, a *Album) {
	err := templates.ExecuteTemplate(w, tmpl+".html", a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	fmt.Println("Record Holder")
	fmt.Println("server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}