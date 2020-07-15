package main

import (
	"database/sql"
	"io"
	"net/http"
	"regexp"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var err error

type user struct {
	ID       int
	Username string
	Role     string
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
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/private", private)
	http.HandleFunc("/register", register)
	http.HandleFunc("/remove/", remove)
	http.HandleFunc("/update", updateUser)
	http.HandleFunc("/users", showUsers)
	http.HandleFunc("/loggedinusers", showLoggedInUsers)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":80", nil)
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
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

func private(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("goSession")

	if err != nil {
		http.Redirect(w, req, "/login", http.StatusForbidden)
		return
	}

	un := activeSessions[c.Value]

	if !isLoggedInAdmin(un) {
		http.Redirect(w, req, "/login", http.StatusForbidden)
		return
	}
	usr := loggedInUsers[un]
	tpl.ExecuteTemplate(w, "private.html", usr)
}

func updateUser(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("goSession")

	if err != nil {
		http.Redirect(w, req, "/login", http.StatusForbidden)
		return
	}

	un := activeSessions[c.Value]

	if !isLoggedInAdmin(un) {
		http.Redirect(w, req, "/login", http.StatusForbidden)
		return
	}

	if req.Method == "POST" {
		oldUsername := req.FormValue("findUser")
		newUsername := req.FormValue("username")
		pwrd := req.FormValue("newpassword")
		newPassword, err := bcrypt.GenerateFromPassword([]byte(pwrd), bcrypt.MinCost)
		check(err, "bcrypt.GenerateFromPassword")

		update(newUsername, oldUsername, pwrd, newPassword)
		delete(dbUsers, oldUsername)
		http.Redirect(w, req, "/users", http.StatusSeeOther)
		return
	}
	getUsers()
	tpl.ExecuteTemplate(w, "update.html", dbUsers)
}

func remove(w http.ResponseWriter, req *http.Request) {
	q := req.URL.String()

	rex := regexp.MustCompile("/remove/")
	userID := rex.ReplaceAllString(q, "")
	deleteUser(userID)

	http.Redirect(w, req, "/users", http.StatusSeeOther)
}
