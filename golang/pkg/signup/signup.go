package signup

import (
	"net/http"
	"test/pkg/tools"
)

func SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	tools.RenderTemplates(w, "../../pkg/signup/templates/signup.html", nil)
}
