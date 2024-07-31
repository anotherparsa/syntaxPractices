package signup

import (
	"net/http"
	"os"
	"test/tools"
)

func SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	dir, _ := os.Getwd()
	tools.RenderTemplates(w, dir+"/signup/signup.html", nil)
}
