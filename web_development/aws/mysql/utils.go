package main

import (
	"log"
	"net/http"
)

// ErrorMessages is a collection of error messages.
type ErrorMessages struct {
	usernameTaken string
	nouser        string
}

var errMssgs = ErrorMessages{
	"Username is taken. Please try again.",
	"The username and/or password were incorrect. Please try again.",
}

func check(err error, from string) {
	if err != nil {
		log.Fatalf("Error from %s: %v", from, err)
	}
}

func isLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := activeSessions[c.Value]
	_, ok := loggedInUsers[un]
	return ok
}
