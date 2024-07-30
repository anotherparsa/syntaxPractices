package home

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type datatosend struct {
	D1 string
	D2 string
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Please go back to sign up page and try again")
	} else {
		dir, _ := os.Getwd()
		t, _ := template.ParseFiles(dir + "/home/home.html")
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		d := datatosend{D1: username, D2: password}
		t.Execute(w, d)
	}
}
