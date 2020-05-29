// Pass data through form method POST
// NOTE: By default a form submit has the method GET. You can make this explicit or leave it implicit.
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
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST">
	<input type="text" name="q"/>
	<input type="submit" />
	</form>
	<br />
	`+v)
}
