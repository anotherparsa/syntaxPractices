package signup

import (
	"net/http"
	"test/pkg/tools"
)

type datatosend struct {
	D1 string
}

var Csrf_token string

func SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	Csrf_token = tools.GenerateCSRFT()
	d := datatosend{D1: Csrf_token}
	tools.RenderTemplates(w, "../../pkg/signup/templates/signup.html", d)
}
