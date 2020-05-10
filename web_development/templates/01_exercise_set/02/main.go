package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./*.gohtml"))
}

func main() {
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("Could not create the file: %v", err)
	}
	defer nf.Close()
	err = tpl.ExecuteTemplate(nf, tpl.Name(), caliHotels)
	if err != nil {
		log.Fatalf("Could not execute the template: %v", err)
	}
}
