package home

import (
	"fmt"
	"net/http"
	"test/pkg/signup"
	"test/pkg/tools"
	"time"
)

type datatosend struct {
	D1 string
	D2 string
	D3 string
	D4 string
	D5 string
	D6 string
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		time.Sleep(time.Second * 5)
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else {
		real_csrf := signup.Csrf_token
		r.ParseForm()
		username := r.Form.Get("username")
		firsname := r.Form.Get("firstname")
		lastname := r.Form.Get("lastname")
		password := r.Form.Get("password")
		sent_csrf_token := r.Form.Get("csrf-token")
		if real_csrf != sent_csrf_token {
			d := datatosend{D3: real_csrf, D4: sent_csrf_token}
			tools.RenderTemplates(w, "../../pkg/home/templates/failedHome.html", d)
		} else {
			user := tools.User{Username: username, Firstname: firsname, Lastname: lastname, Password: password}
			userid := tools.GenerateUUID()
			sessionid := tools.GenerateUUID()
			tools.SetCookie(w, r, userid, sessionid)
			tools.Dbusers[userid] = user
			tools.Dbsession[sessionid] = userid
			d := datatosend{D1: sessionid, D2: userid, D3: user.Username, D4: user.Firstname, D5: user.Lastname, D6: user.Password}
			tools.RenderTemplates(w, "../../pkg/home/templates/home.html", d)
			fmt.Println(tools.Dbsession)
			fmt.Println(tools.Dbusers)
		}

	}
}

func SetupFileserver() {
	fileserver := http.FileServer(http.Dir("../../pkg/home/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
}
