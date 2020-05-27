// EXERCISE 02 (follows on from EXERCISE 01)

// (EXERCISE 01)
// ListenAndServe on port ":8080" using the default ServeMux.
// Use HandleFunc to add the following routes to the default ServeMux:
// "/" "/dog/" "/me/
// Add a func for each of the routes.
// Have the "/me/" route print out your name.

// (EXERCISE 02)
// Parse and serve a template with data you pass into it.

package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func thingOne(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, req.Method, req.URL.Path)
}

func thingTwo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, req.Method, req.URL.Path)
}

func thingThree(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Fprintln(w, req.Method, req.URL.Path, req.Form.Get("fname"))
}

func thingFour(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", req.Method)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", thingOne)
	http.HandleFunc("/dog/", thingTwo)
	http.HandleFunc("/me", thingThree)
	http.HandleFunc("/tf", thingFour)

	http.ListenAndServe(":8080", nil)
}
