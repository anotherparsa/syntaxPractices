package tools

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
)

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
