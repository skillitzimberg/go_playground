// Using templates to build a web page . . .
// Pass data using form method POST.
// Use a struct that maps to the form values.
// Pass this data as a value (e.g. p := person{... <form values> ...}) to a template.
package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

// Person defines a generic person.
type Person struct {
	FirstName  string
	LastName   string
	IsLicensed bool
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	f := req.FormValue("fName")
	l := req.FormValue("lName")
	il := req.FormValue("isLicensed") == "on"

	err := tpl.ExecuteTemplate(w, "index.gohtml", Person{f, l, il})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
