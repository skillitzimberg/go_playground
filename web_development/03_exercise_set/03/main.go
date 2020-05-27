// Serve the files in the "starting-files" folder
// To get your images to serve, use only this:
// 	fs := http.FileServer(http.Dir("public"))
// Hint: look to see what type FileServer returns, then think it through.
package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/images/", fs)
	http.Handle("/pics/", fs)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
