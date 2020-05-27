// EXERCISE 03 (follows on from EXERCISE 02)

// (EXERCISE 01)
// ListenAndServe on port ":8080" using the default ServeMux.
// Use HandleFunc to add the following routes to the default ServeMux:
// "/" "/dog/" "/me/
// Add a func for each of the routes.
// Have the "/me/" route print out your name.

// (EXERCISE 02)
// Parse and serve a template with data you pass into it.

// (EXERCISE 03)
// func main uses http.Handle instead of http.HandleFunc
// Constraint: Do not change anything outside of func main from exercise 02.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func thingOne(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Thing One says:", req.Method, req.URL.Path)
}

func thingTwo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Thing Two says:", req.Method, req.URL.Path)
}

func thingThree(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Fprintln(w, "Thing Three says,", "Method:", req.Method, "Path:", req.URL.Path, "Query params fname:", req.Form.Get("fname"))
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

	http.Handle("/", http.HandlerFunc(thingOne))
	http.Handle("/dog/", http.HandlerFunc(thingTwo))
	http.Handle("/me", http.HandlerFunc(thingThree))
	http.Handle("/tf", http.HandlerFunc(thingFour))

	http.ListenAndServe(":8080", nil)
}
