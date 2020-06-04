package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

// map[sessionId]username
var dbSessions = map[string]string{}

// map[username]user
var dbUsers = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	if !isLoggedIn(u.UserName) {
		http.Redirect(w, req, "/signup", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	if isLoggedIn(u.UserName) {
		http.Redirect(w, req, "/bar", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("UserName")
		pwrd := req.FormValue("Password")
		f := req.FormValue("First")
		l := req.FormValue("Last")
		p, err := bcrypt.GenerateFromPassword([]byte(pwrd), bcrypt.MinCost)
		if err != nil {
			log.Fatal(err)
		}
		u = user{
			un, p, f, l,
		}

		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:     "session",
			Value:    sID.String(),
			HttpOnly: true,
		}

		dbSessions[c.Value] = u.UserName
		dbUsers[u.UserName] = u
		http.SetCookie(w, c)

		http.Redirect(w, req, "/bar", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func logout(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	delete(dbUsers, u.UserName)
	fmt.Println(dbUsers)

	http.Redirect(w, req, "/", http.StatusSeeOther)
	return
}
