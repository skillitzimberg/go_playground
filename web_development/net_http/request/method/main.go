package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type formTemplate struct{}

func (ft formTemplate) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(tpl.ParseFiles("index.gohtml"))
}

func main() {
	var ft formTemplate
	err := http.ListenAndServe(":8080", ft)
	if err != nil {
		log.Fatal(err)
	}
}
