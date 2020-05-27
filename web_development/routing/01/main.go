// ListenAndServe on port ":8080" using the default ServeMux.

// Use HandleFunc to add the following routes to the default ServeMux:

// "/" "/dog/" "/me/

// Add a func for each of the routes.

// Have the "/me/" route print out your name.
package main

import (
	"fmt"
	"net/http"
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

func main() {
	http.HandleFunc("/", thingOne)
	http.HandleFunc("/dog/", thingTwo)
	http.HandleFunc("/me", thingThree)

	http.ListenAndServe(":8080", nil)
}
