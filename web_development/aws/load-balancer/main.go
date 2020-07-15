package main

import (
	"database/sql"
	"html/template"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var err error

var connectionString = "admin:sbmzggX0@tcp(crud-rds-db.cbcbeyzcudgn.us-west-2.rds.amazonaws.com:3306)/CRUD_RDS_DB?charset=utf8"
var pool *sql.DB // Database connection pool.

type user struct {
	ID       int
	Username string
}

type pageData struct {
	Instance string
	Users    map[string]user
}

// map[username]user
var dbUsers = map[string]user{}

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
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/users", users)
	http.HandleFunc("/instance", instance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	s := "Hello from AWS"
	s += getInstanceID()
	tpl.ExecuteTemplate(w, "index.html", s)
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

func users(w http.ResponseWriter, req *http.Request) {
	rows, err := pool.Query("SELECT id, username FROM users")
	check(err, "pool.Query")
	defer rows.Close()

	var id int
	var username string
	instanceID := getInstanceID()

	for rows.Next() {
		err = rows.Scan(&id, &username)
		check(err, "rows.Scan")
		dbUsers[username] = user{id, username}
	}
	pg := pageData{instanceID, dbUsers}
	tpl.ExecuteTemplate(w, "users.html", pg)
}

func instance(w http.ResponseWriter, req *http.Request) {
	instanceID := getInstanceID()
	io.WriteString(w, instanceID)
}

func getInstanceID() string {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		check(err, "getInstanceID")
		return err.Error()
	}

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	if err != nil {
		check(err, "Body.Read")
		return err.Error()
	}
	defer resp.Body.Close()
	return string(bs)
}

func check(err error, from string) {
	if err != nil {
		log.Fatalf("Error from %s: %v", from, err)
	}
}
