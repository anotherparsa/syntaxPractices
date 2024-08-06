package tools

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Username  string
	Firstname string
	Lastname  string
	Password  string
}

var Dbusers = map[string]User{}     //user id -> user
var Dbsession = map[string]string{} //session id -> user id

type datatosend struct {
	D1 string
}

/*func RenderTemplates(w http.ResponseWriter, templateAddress string, data any) {
	parsedTemplate, err := template.ParseFiles(templateAddress, "../../pkg/baseTemplate/baseTemplate.html")
	if err != nil {
		fmt.Println(err)
	} else {
		parsedTemplate.Execute(w, data)
	}
}*/

var templateCache = make(map[string]*template.Template)

func RenderTemplates(w http.ResponseWriter, templateAddress string, data any) {
	var template *template.Template
	var err error

	//check if we already have a parsed template
	_, inMap := templateCache[templateAddress]
	if !inMap {
		//need to parse from scratch
		fmt.Println("It wasn't in the cache and now we're creating and adding it to the map")
		err = createTemplateCache(templateAddress)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		//we have a parsed template
		fmt.Println("Using parsed template")
	}

	template = templateCache[templateAddress]
	err = template.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func createTemplateCache(templateAddress string) error {
	templates := []string{
		templateAddress,
		"../../pkg/baseTemplate/baseTemplate.html",
	}
	//parse the template
	parsedTemplate, err := template.ParseFiles(templates[0], templates[1])

	if err != nil {
		return err
	} else {
		//add parsed template to the cache map
		templateCache[templateAddress] = parsedTemplate
		return nil
	}

}

func GenerateCSRFTPageHandler(w http.ResponseWriter, r *http.Request) {
	d := datatosend{D1: GenerateCSRFT()}
	RenderTemplates(w, "../../pkg/tools/templates/tools.html", d)
}

func GenerateCSRFT() string {
	bytesSlices := make([]byte, 32)
	_, err := rand.Read(bytesSlices)
	if err != nil {
		fmt.Println(err)
	}

	return base64.StdEncoding.EncodeToString(bytesSlices)
}

func GenerateUUID() string {
	byteSlices := make([]byte, 33)
	_, err := rand.Read(byteSlices)
	if err != nil {
		fmt.Println(err)
	}

	return base64.StdEncoding.EncodeToString(byteSlices[:33])
}

func SetCookie(w http.ResponseWriter, r *http.Request, user_id string, session_id string) {
	http.SetCookie(w, &http.Cookie{Name: "session_id", Value: user_id})
}

func SetCookieTwo(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "Name", Value: "Value"})

}

func ReadCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	fmt.Fprintf(w, "Your cooki is %v", cookie)

}

func ExpireCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	fmt.Fprintf(w, "The cookie has been expired")
}
