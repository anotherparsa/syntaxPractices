package home

import (
	"fmt"
	"net/http"
	"os"
	"test/tools"
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
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		d := datatosend{D1: username, D2: password}
		tools.RenderTemplates(w, dir+"/home/home.html", d)
	}
}
