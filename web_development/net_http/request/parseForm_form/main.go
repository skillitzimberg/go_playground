package main

import (
	"log"
	"net/http"
	"text/template"
)

type formTemplate struct{}

var tpl *template.Template

func (ft formTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var ft formTemplate
	err := http.ListenAndServe(":8080", ft)
	if err != nil {
		log.Fatal(err)
	}
}
