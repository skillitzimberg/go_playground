package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("views/*.html"))
}

func serveCSS() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css", fs))
}

func main() {
	serveCSS()
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/projects", projects)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("localhost:8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func about(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "about.html", nil)
}

func contact(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "contact.html", nil)
}

func projects(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "projects.html", nil)
}
