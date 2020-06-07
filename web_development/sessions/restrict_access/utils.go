package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

const session = "session"

func getUser(req *http.Request) user {
	u := user{}
	c := getCookie(req)

	un, ok := dbSessions[c.Value]
	if ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c := getCookie(req)
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}

func newUUID(w http.ResponseWriter) uuid.UUID {
	sID, err := uuid.NewV4()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return sID
}

func getCookie(req *http.Request) *http.Cookie {
	c, err := req.Cookie(session)
	if err != nil {
		return &http.Cookie{
			Name:  session,
			Value: "",
		}
	}
	return c
}
