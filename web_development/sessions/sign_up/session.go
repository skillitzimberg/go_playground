package main

import (
	"net/http"
)

// Logged in means these three things are true:
// 1. there is a valid cookie
// 2. the cookie has been added to the dbSessions map
// 3. the user has been added to the dbUsers map
func isLoggedIn(username string) bool {
	_, ok := dbUsers[username]
	return ok
}

func getUser(req *http.Request) user {
	var u user
	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	if username, ok := dbSessions[c.Value]; ok {
		u = dbUsers[username]
	}

	return u
}
