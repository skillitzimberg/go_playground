package main

import "fmt"

func getUsers() {
	rows, err := pool.Query("SELECT username, password FROM users")
	check("pool.Query", err)
	defer rows.Close()

	var username, password string

	for rows.Next() {
		err = rows.Scan(&username, &password)
		check("rows.Scan", err)
		users[username] = user{username, []byte("")}
	}
}

func createUser(username string, password []byte) {
	s := fmt.Sprintf(`INSERT INTO users (username, password) VALUES ("%s", "%s");`, username, password)
	stmt, err := pool.Prepare(s)
	check("pool.Prepare", err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check("stmt.Exec", err)
}
