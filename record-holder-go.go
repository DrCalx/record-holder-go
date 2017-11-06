package main

import (
	"fmt"
	"os"
	"html/template"
	"net/http"
)

type Album struct {
	Id string
	Title string
	Artist string
}

var templates = template.Must(template.ParseFiles("views/layout.html", "views/view.html", "views/edit.html"))

func viewHandler(w http.ResponseWriter, r *http.Request) {
	a := &Album{Id: "1", Title: "Please Please Me", Artist: "The Beatles"}
	renderTemplate(w, "view", a)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	a := &Album{Id: "1", Title: "Please Please Me", Artist: "The Beatles"}
	renderTemplate(w, "edit", a)
}

func renderTemplate(w http.ResponseWriter, tmpl string, a *Album) {
	t, _ := template.ParseFiles("views/layout.html", "views/"+tmpl+".html")
	err := t.ExecuteTemplate(w, "layout", a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	fmt.Println("Record Holder")
	fmt.Println("server listening on port", port)
	http.ListenAndServe(":"+port, nil)
}