package main

import (
	"fmt"
	"net/http"
)

func getUsers() {
	rows, err := pool.Query("SELECT username, id FROM users")
	check("pool.Query", err)
	defer rows.Close()

	var username, id string

	for rows.Next() {
		err = rows.Scan(&username, &id)
		check("rows.Scan", err)
		users[username] = user{username, id}
	}
}

func createUser(username string, password string) {
	s := fmt.Sprintf(`INSERT INTO users (username, password) VALUES ("%s", "%s");`, username, password)
	stmt, err := pool.Prepare(s)
	check("pool.Prepare", err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check("stmt.Exec", err)
}

func checkForUser(w http.ResponseWriter, username string, password string) error {
	s := fmt.Sprintf(`SELECT id FROM users WHERE password="%s" LIMIT 1`, string(password))
	r := pool.QueryRow(s)
	var user string
	err := r.Scan(&user)
	return err
}
