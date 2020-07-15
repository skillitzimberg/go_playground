package main

import (
	"database/sql"
	"fmt"
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

// map[username]user
var dbUsers = map[string]user{}

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
	s := "Hello from AWS: Round Three."
	s += getInstanceID()
	io.WriteString(w, s)
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
	s := getInstanceID()
	s += "\nREGISTERED USERS:\n"
	fmt.Println(s)

	for rows.Next() {
		fmt.Println(s)
		err = rows.Scan(&id, &username)
		check(err, "rows.Scan")
		dbUsers[username] = user{id, username}
		s += username + "\n"
	}
	io.WriteString(w, s)
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
