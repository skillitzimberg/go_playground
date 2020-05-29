// Pass a query through the url.
package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(w, "You asked for "+v+".")
}
