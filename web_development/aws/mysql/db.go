package main

import (
	"database/sql"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

var connectionString = "admin:sbmzggX0@tcp(database-1.cbcbeyzcudgn.us-west-2.rds.amazonaws.com:3306)/gowebdev?charset=utf8"
var pool *sql.DB // Database connection pool.

func getUsers() {
	rows, err := pool.Query("SELECT id, role, username FROM users")
	check(err, "pool.Query")
	defer rows.Close()

	var id int
	var username string
	var role string

	for rows.Next() {
		err = rows.Scan(&id, &role, &username)
		check(err, "rows.Scan")
		dbUsers[username] = user{id, username, role}
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

func saveNewUser(username string, password string) {
	hashedPwrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	check(err, "bcrypt.GenerateFromPassword")
	s := fmt.Sprintf(`INSERT INTO users (username, password, role) VALUES ("%s", "%s", "%s");`, username, hashedPwrd, " ")
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

func update(newUsername string, oldUsername string, pwrd string, newHashedPassword []byte) {
	var s string
	if pwrd == "" {
		s = fmt.Sprintf(`UPDATE users SET username="%s" WHERE username="%s"`, newUsername, oldUsername)
	} else if newUsername == "" {
		s = fmt.Sprintf(`UPDATE users SET password="%s" WHERE username="%s"`, newHashedPassword, oldUsername)
	} else {
		s = fmt.Sprintf(`UPDATE users SET username="%s", password="%s" WHERE username="%s"`, newUsername, newHashedPassword, oldUsername)
	}

	stmt, err := pool.Prepare(s)
	check(err, "pool.Prepare")
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err, "stmt.Exec")
}

func deleteUser(userID string) {
	fmt.Println(userID)
	s := fmt.Sprintf(`DELETE FROM users WHERE id="%s"`, userID)
	stmt, err := pool.Prepare(s)
	check(err, "pool.Prepare")
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err, "stmt.Exec")

	id, _ := strconv.Atoi(userID)
	un := findUsernameByID(dbUsers, id)
	delete(dbUsers, un)
	fmt.Println(dbUsers)
}
