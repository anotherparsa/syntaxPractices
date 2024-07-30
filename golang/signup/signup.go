package signup

import (
	"html/template"
	"net/http"
	"os"
)

func SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	dir, _ := os.Getwd()
	t, _ := template.ParseFiles(dir + "/signup/signup.html")
	t.Execute(w, nil)
}
