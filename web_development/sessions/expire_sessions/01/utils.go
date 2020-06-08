package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	_, ok = dbUsers[s.username]
	return ok
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}

func encryptPassword(w http.ResponseWriter, password string) []byte {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return []byte{}
	}
	return p
}

func getUser(w http.ResponseWriter, req *http.Request) user {
	// create a zero value user
	var u user
	// get the current cookie/session
	c := getCookie(w, req)
	http.SetCookie(w, c)

	// if the session exists return the found user
	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.username]
	}
	// otherwise return the empty user
	return u
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	// get the current cookie/session
	c, err := req.Cookie("session")

	// there isn't one so return an new cookie/session
	if err != nil {
		sID, err := uuid.NewV4()
		handleInternalServerError(w, err)

		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}

	c.MaxAge = sessionLength
	return c
}

func newUUID(w http.ResponseWriter) uuid.UUID {
	id, err := uuid.NewV4()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return id
}

func handleInternalServerError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func passwordMatches(userPassword []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(userPassword, []byte(password))
	if err == nil {
		return true
	}
	return false
}

func showSessions() {
	fmt.Println("******")
	for k := range dbSessions {
		fmt.Println(k)
	}
	fmt.Println("")
}
