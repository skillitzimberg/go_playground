package main

import (
	"fmt"
	"net/http"
)

type myType string

func (mt myType) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, mt)
}

func main() {
	dog := "I'm a dog, Dan."

	http.ListenAndServe(":8080", myType(dog))
}
