package views

import (
	//"golang.org/x/net/html"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	if _, ok := req.Cookie("sessionID"); ok != nil {
		http.Redirect(w, req, "/login", http.StatusTemporaryRedirect)
	}
	tpl := template.Must(template.ParseFiles("html/index.html"))
	_ = tpl.Execute(w, nil)
}

func Insert(w http.ResponseWriter, _ *http.Request) {
	tpl := template.Must(template.ParseFiles("html/insert.html"))
	_ = tpl.Execute(w, nil)
}
