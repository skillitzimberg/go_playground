package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
)

var err error

type user struct {
	ID       int
	Username string
}

// map[username]user
var dbUsers = map[string]user{}

// map[username]user
var loggedInUsers = map[string]user{}

// map[seesionID]username
var activeSessions = map[string]string{}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	pool, err = sql.Open("mysql", connectionString)
	check(err, "sql.Open")
	defer pool.Close()

	err = pool.Ping()
	check(err, "pool.Ping")

	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/register", register)
	http.HandleFunc("/users", showUsers)
	http.HandleFunc("/loggedinusers", showLoggedInUsers)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe("localhost:8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func showUsers(w http.ResponseWriter, req *http.Request) {
	getUsers()
	tpl.ExecuteTemplate(w, "users.html", dbUsers)
}

func showLoggedInUsers(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "users.html", loggedInUsers)
}

func register(w http.ResponseWriter, req *http.Request) {
	getUsers()

	if req.Method == "POST" {
		un := req.FormValue("username")
		pwrd := req.FormValue("password")

		if _, ok := dbUsers[un]; ok {
			tpl.ExecuteTemplate(w, "register.html", errMssgs.usernameTaken)
			return
		}

		saveNewUser(un, pwrd)
		usr := getUserFromDB(un)
		dbUsers[un] = usr
		http.Redirect(w, req, "/users", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "register.html", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	getUsers()

	if isLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	if req.Method == "POST" {
		un := req.FormValue("username")
		pwrd := req.FormValue("password")

		if !isRegistered(un, pwrd) {
			tpl.ExecuteTemplate(w, "login.html", errMssgs.nouser)
			return
		}

		sID, err := uuid.NewV4()
		check(err, "uuid.NewV4")
		c := &http.Cookie{
			Name:  "goSession",
			Value: sID.String(),
		}
		activeSessions[c.Value] = un
		loggedInUsers[un] = dbUsers[un]
		http.SetCookie(w, c)

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.html", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !isLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	c, err := req.Cookie("goSession")
	check(err, "goSession")

	un := activeSessions[c.Value]
	delete(loggedInUsers, un)
	delete(activeSessions, c.Value)
	c.MaxAge = -1
	http.SetCookie(w, c)

	tpl.ExecuteTemplate(w, "users.html", loggedInUsers)
}
