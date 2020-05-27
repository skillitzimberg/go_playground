// ListenAndServe on port 8080 of localhost
// For the default route "/" Have a func called "foo" which writes to the response "foo ran"

// For the route "/dog/" Have a func called "dog" which parses a template called "dog.gohtml" and writes to the response "

// This is from dog
// " and also shows a picture of a dog when the template is executed.
// Use "http.ServeFile" to serve the file "dog.jpeg"
package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/pittbull.jpeg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>Foo ran.</h1>")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("dog.gohtml"))

	err := tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	io.WriteString(w, "<h1>This is from dog.</h1>")
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "pittbull.jpeg")
}
