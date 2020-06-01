// Allow user to set a cookie
// Allow user to read a cookie
// Allow a user to expire a cookie
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<h1><a href="/set">Set</a></h1>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "session",
			Value: "I'm a cookie!",
		}
	}

	http.SetCookie(w, c)
	fmt.Fprintln(w, `<h1><a href="/read">Read</a></h1>`)

}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, `<p>Cookie value: %s</p><h1><a href="/expire">Expire</a></h1>`, c.Value)
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	c.MaxAge = -1
	http.SetCookie(w, c)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}
