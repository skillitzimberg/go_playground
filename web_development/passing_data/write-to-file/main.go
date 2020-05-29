// Allow a user to upload a file.
// Read the file contents.
// Write the file contents in the browser.
// Write the file context to a new file
package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
		f, h, err := req.FormFile("f")
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

		// create a new file
		cf, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cf.Close()

		// write original file contents to the new file
		_, err = cf.Write(fbs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// write file contents to the browser
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(w, "index.gohtml", s)
	if err != nil {
		log.Fatalln("Template could not be executed:", err)
	}
}

func handleInternalServerError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
