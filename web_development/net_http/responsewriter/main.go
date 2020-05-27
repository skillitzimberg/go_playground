package main

import (
	"fmt"
	"log"
	"net/http"
)

type sommat struct{}

func (s sommat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Scott-Key", "Something Scott wanted in a header")
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>The Content Requested.</h1>")

}

func main() {
	var s sommat
	err := http.ListenAndServe(":8080", s)
	if err != nil {
		log.Fatal(err)
	}
}
