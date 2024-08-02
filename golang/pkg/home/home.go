package home

import (
	"fmt"
	"net/http"
	"test/pkg/signup"
	"test/pkg/tools"
)

type datatosend struct {
	D1 string
	D2 string
	D3 string
	D4 string
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Please go back to sign up page and try again")
	} else {
		real_csrf := signup.Get_generated_csrf_token()
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		sent_csrf_token := r.Form.Get("csrf-token")
		if real_csrf != sent_csrf_token {
			d := datatosend{D3: real_csrf, D4: sent_csrf_token}
			tools.RenderTemplates(w, "../../pkg/home/templates/failedHome.html", d)
		} else {
			d := datatosend{D1: username, D2: password, D3: real_csrf, D4: sent_csrf_token}
			tools.RenderTemplates(w, "../../pkg/home/templates/home.html", d)
		}

	}
}
