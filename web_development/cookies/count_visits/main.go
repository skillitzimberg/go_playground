// Using cookies, track how many times a user has been to your domain.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")

	if err != nil {
		c = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path:  "/",
		}
	}

	visits, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatal(err)
	}
	visits++
	c.Value = strconv.Itoa(visits)

	http.SetCookie(w, c)
	fmt.Fprintf(w, "You have been here %s times.", c.Value)
}
