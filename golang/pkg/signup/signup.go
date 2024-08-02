package signup

import (
	"net/http"
	"test/pkg/tools"
)

type datatosend struct{
	D1 string
}

func SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	csrf_token := tools.GenerateCSRFT()
	d := datatosend{D1: csrf_token}
	tools.RenderTemplates(w, "../../pkg/signup/templates/signup.html", d)
}
