package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var connectionString = "admin:sbmzggX0@tcp(database-1.cbcbeyzcudgn.us-west-2.rds.amazonaws.com:3306)/gowebdev?charset=utf8"
var pool *sql.DB // Database connection pool.
var err error

type user struct {
	Username string
	Password []byte
}

var users = map[string]user{} // [username]user

// ErrorMessages is a library of error messages.
type ErrorMessages struct {
	usernameTaken string
}

var errMssgs = ErrorMessages{"Username is taken. Please try again."}

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
	getUsers()
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func showUsers(w http.ResponseWriter, req *http.Request) {
	getUsers()
	tpl.ExecuteTemplate(w, "users.html", users)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		un := req.FormValue("username")
		pwrd := req.FormValue("password")

		if _, ok := users[un]; ok {
			tpl.ExecuteTemplate(w, "signup.html", errMssgs.usernameTaken)
			return
		}

		p, err := bcrypt.GenerateFromPassword([]byte(pwrd), bcrypt.MinCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		createUser(un, p)

		http.Redirect(w, req, "/users", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "signup.html", nil)
}
