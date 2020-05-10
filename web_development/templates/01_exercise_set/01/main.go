package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("Error creating a file %v", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, years)
	if err != nil {
		log.Fatalln(err)
	}
}
