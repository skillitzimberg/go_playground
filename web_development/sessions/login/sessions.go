package main

import (
	"net/http"
)

func getUser(req *http.Request) user {
	u := user{}

	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	return u
}

func isLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
