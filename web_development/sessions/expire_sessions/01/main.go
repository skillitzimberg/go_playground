package main

import (
	"net/http"
	"text/template"
	"time"
)

type user struct {
	Username  string
	Password  []byte
	FirstName string
	LastName  string
	Role      string
}

type session struct {
	username     string
	lastActivity time.Time
}

const sessionLength int = 30

var dbUsers = map[string]user{}       // username: user
var dbSessions = map[string]session{} //sessionID: session
var dbSessionsCleaned time.Time
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/armory", armory)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	showSessions()
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	u := getUser(w, req)
	showSessions()
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func armory(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	u := getUser(w, req)
	if u.Role != "007" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "armory.gohtml", u)
}

func register(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var u user
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		pwrd := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username is already taken.", http.StatusForbidden)
			return
		}

		p := encryptPassword(w, pwrd)
		u = user{un, p, f, l, r}
		dbUsers[un] = u

		http.Redirect(w, req, "/login", http.StatusSeeOther)
	}
	showSessions()
	tpl.ExecuteTemplate(w, "register.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")

		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username not found.", http.StatusForbidden)
			return
		}

		if !passwordMatches(u.Password, p) {
			http.Error(w, "Password not found.", http.StatusForbidden)
			return
		}

		c := getCookie(w, req)
		c.MaxAge = sessionLength
		http.SetCookie(w, c)

		dbSessions[c.Value] = session{un, time.Now()}

		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	showSessions()
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	c := getCookie(w, req)
	delete(dbSessions, c.Value)
	c.MaxAge = -1
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
}
