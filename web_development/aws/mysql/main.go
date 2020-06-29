package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
)

var connectionString = "admin:sbmzggX0@tcp(database-1.cbcbeyzcudgn.us-west-2.rds.amazonaws.com:3306)/gowebdev?charset=utf8"
var pool *sql.DB // Database connection pool.
var err error

type user struct {
	ID       int
	Username string
}

// map[username]user
var users = map[string]user{}

// map[seesionID]username
var sessions = map[string]string{}

// ErrorMessages is a collection of error messages.
type ErrorMessages struct {
	usernameTaken string
	nouser        string
}

var errMssgs = ErrorMessages{
	"Username is taken. Please try again.",
	"The username and/or password were incorrect. Please try again.",
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	pool, err = sql.Open("mysql", connectionString)
	check("sql.Open", err)
	defer pool.Close()

	err = pool.Ping()
	check("pool.Ping", err)

	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/users", showUsers)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe("localhost:8080", nil)
}

func check(from string, err error) {
	if err != nil {
		log.Fatalf("Error from %s: %v", from, err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func showUsers(w http.ResponseWriter, req *http.Request) {
	getUsers()
	tpl.ExecuteTemplate(w, "users.html", users)
}

func signup(w http.ResponseWriter, req *http.Request) {
	getUsers()

	if req.Method == "POST" {
		un := req.FormValue("username")
		pwrd := req.FormValue("password")

		if _, ok := users[un]; ok {
			tpl.ExecuteTemplate(w, "signup.html", errMssgs.usernameTaken)
			return
		}

		createUser(un, pwrd)
		usr := getUser(un)
		users[un] = usr
		http.Redirect(w, req, "/users", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func login(w http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		un := req.FormValue("username")
		pwrd := req.FormValue("password")

		if !isRegistered(un, pwrd) {
			tpl.ExecuteTemplate(w, "login.html", errMssgs.nouser)
			return
		}

		sID, err := uuid.NewV4()
		check("uuid.NewV4", err)
		c := &http.Cookie{
			Name:  "goSession",
			Value: sID.String(),
		}
		sessions[c.Value] = un
		http.SetCookie(w, c)

		http.Redirect(w, req, "index.html", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.html", nil)
}
