// Create User (struct), Users([]User), sessions (map[uuid]UserName).
// Create a template form to get user info from the user (UserName, First, Last, etc.) and a link to another view.
// Display the user's info in both views.
// Create a session.
// Use these tools to retrieve and display user info.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

type user struct {
	UserName string
	First    string
	Last     string
}

var dbUsers = map[string]user{}

var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		sID, err := uuid.NewV4()
		if err != nil {
			log.Fatalln(err)
		}
		c = &http.Cookie{
			Name:     "session",
			Value:    sID.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, c)
	}

	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("UserName")
		f := req.FormValue("First")
		l := req.FormValue("Last")
		u = user{un, f, l}

		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		fmt.Println("no cookie/session.")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSessions[c.Value]
	if !ok {
		fmt.Println("User does not exist.")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}


