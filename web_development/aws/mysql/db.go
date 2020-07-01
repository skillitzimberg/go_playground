package main

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var connectionString = "admin:sbmzggX0@tcp(database-1.cbcbeyzcudgn.us-west-2.rds.amazonaws.com:3306)/gowebdev?charset=utf8"
var pool *sql.DB // Database connection pool.

func getUsers() {
	rows, err := pool.Query("SELECT id, username FROM users")
	check(err, "pool.Query")
	defer rows.Close()

	var id int
	var username string

	for rows.Next() {
		err = rows.Scan(&id, &username)
		check(err, "rows.Scan")
		dbUsers[username] = user{id, username}
	}
}

func getUserFromDB(username string) user {
	var user = user{}
	s := fmt.Sprintf(`SELECT id, username FROM users WHERE username="%s" LIMIT 1`, username)
	r := pool.QueryRow(s)
	err := r.Scan(&user.ID, &user.Username)
	check(err, "r.Scan")
	return user
}

func getLoggedInUser(username string) user {
	return loggedInUsers[username]
}

func saveNewUser(username string, password string) {
	hashedPwrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	check(err, "bcrypt.GenerateFromPassword")
	s := fmt.Sprintf(`INSERT INTO users (username, password) VALUES ("%s", "%s");`, username, hashedPwrd)
	stmt, err := pool.Prepare(s)
	check(err, "pool.Prepare")
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err, "stmt.Exec")
}

func isRegistered(username string, password string) bool {
	var dbPwrd []byte
	s := fmt.Sprintf(`SELECT password FROM users WHERE username="%s" LIMIT 1`, username)
	r := pool.QueryRow(s)
	err := r.Scan(&dbPwrd)
	err = bcrypt.CompareHashAndPassword(dbPwrd, []byte(password))
	return err == nil
}
