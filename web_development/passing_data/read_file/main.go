// Allow a user to upload a file.
// Read the file contents.
// Write the file contents in the browser.
package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", fileHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func fileHandler(w http.ResponseWriter, req *http.Request) {
	var s string
	if req.Method == http.MethodPost {
		// open the file
		f, _, err := req.FormFile("f")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// read the file
		fbs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(fbs)
	}

	// write the file to the browser
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(w, "index.gohtml", s)
	if err != nil {
		log.Fatalln("Template could not be executed:", err)
	}
}
