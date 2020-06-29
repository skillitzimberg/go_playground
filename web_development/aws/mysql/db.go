package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func getUsers() {
	rows, err := pool.Query("SELECT id, username FROM users")
	check("pool.Query", err)
	defer rows.Close()

	var id int
	var username string

	for rows.Next() {
		err = rows.Scan(&id, &username)
		check("rows.Scan", err)
		users[username] = user{id, username}
	}
}

func getUser(username string) user {
	var user = user{}
	s := fmt.Sprintf(`SELECT id, username FROM users WHERE username="%s" LIMIT 1`, username)
	r := pool.QueryRow(s)
	err := r.Scan(&user.ID, &user.Username)
	check("r.Scan", err)
	return user
}

func createUser(username string, password string) {
	hashedPwrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	check("bcrypt.GenerateFromPassword", err)
	s := fmt.Sprintf(`INSERT INTO users (username, password) VALUES ("%s", "%s");`, username, hashedPwrd)
	stmt, err := pool.Prepare(s)
	check("pool.Prepare", err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check("stmt.Exec", err)
}

func isRegistered(username string, password string) bool {
	var dbPwrd []byte
	s := fmt.Sprintf(`SELECT password FROM users WHERE username="%s" LIMIT 1`, username)
	r := pool.QueryRow(s)
	err := r.Scan(&dbPwrd)
	err = bcrypt.CompareHashAndPassword(dbPwrd, []byte(password))
	return err == nil
}
