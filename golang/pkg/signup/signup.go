package signup

import (
	"net/http"
	"test/pkg/tools"
)

type datatosend struct {
	D1 string
}

var csrf_token string

func SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	csrf_token = tools.GenerateCSRFT()
	d := datatosend{D1: csrf_token}
	tools.RenderTemplates(w, "../../pkg/signup/templates/signup.html", d)
}

func Get_generated_csrf_token() string {
	return csrf_token
}
