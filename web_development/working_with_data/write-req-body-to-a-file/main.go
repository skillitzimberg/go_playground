// Take user input from a form
// Create a new file on the server
// Write the form contents to the file
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", handleForm)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handleForm(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handleForm")
	if req.Method == http.MethodPost {
		// get the req body data as a slice of bytes
		fmt.Println(req.Body)
		fmt.Println("get the req body")

		bs := make([]byte, req.ContentLength)
		n, err := req.Body.Read(bs)
		if err != nil {
			fmt.Println(n, err)
		}

		// create new file (hint: use os pkg)
		fmt.Println("create new file")

		nf, err := os.Create(filepath.Join("./user/", "request-body.txt"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer nf.Close()

		// write req body data to the file
		fmt.Println("write req body data")
		_, err = nf.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("Could not execute the template:", err)
	}
}
