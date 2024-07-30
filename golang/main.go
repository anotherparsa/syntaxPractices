package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/signup", signupPageHander)
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
