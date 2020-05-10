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
		log.Fatalln("Failed to create a file.", err)
	}
	defer nf.Close()

	err = tpl.ExecuteTemplate(nf, tpl.Name(), menu)
	if err != nil {
		log.Fatalln("Failed to execute the template.", err)
	}

}
