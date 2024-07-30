package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type datatosend struct {
	D1 string
	D2 string
}

func main() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/signup", signupPageHander)
	http.HandleFunc("/signupprocess", signupProcessorHandler)
	fmt.Println("Runnin server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello This is the home page")
}

func signupPageHander(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func signupProcessorHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	d := datatosend{D1: username, D2: password}
	t, _ := template.ParseFiles("home.html")
	t.Execute(w, d)

}
