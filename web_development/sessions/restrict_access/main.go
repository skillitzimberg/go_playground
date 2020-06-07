package main

import (
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username string
	Password []byte
	First    string
	Last     string
	Role     string
}

// map[sessionID]username
var dbSessions = map[string]string{}

// map[username]user
var dbUsers = map[string]user{}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/armory", armory)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/register", register)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func bar(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	u := getUser(req)
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func armory(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := getUser(req)
	if u.Role != "007" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "armory.gohtml", u)
}

func register(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := getUser(req)
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		pwrd := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username is already taken.", http.StatusForbidden)
			return
		}

		p, err := bcrypt.GenerateFromPassword([]byte(pwrd), bcrypt.MinCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		u = user{un, p, f, l, r}
		dbUsers[un] = u

		sID := newUUID(w)
		c := &http.Cookie{
			Name:  session,
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "register.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		pwrd := req.FormValue("password")

		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username.", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(pwrd))
		if err != nil {
			http.Error(w, "Password.", http.StatusForbidden)
			return
		}

		sID := newUUID(w)
		c := &http.Cookie{
			Name:  session,
			Value: sID.String(),
		}
		dbSessions[c.Value] = un
		http.SetCookie(w, c)

		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)

}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	c := getCookie(req)
	delete(dbSessions, c.Value)
	c.MaxAge = -1
	http.SetCookie(w, c)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}
