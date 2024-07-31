package tools

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplates(w http.ResponseWriter, templateAddress string, data any) {
	parsedTemplate, err := template.ParseFiles(templateAddress, "../../pkg/baseTemplate/baseTemplate.html")
	if err != nil {
		fmt.Println(err)
	} else {
		parsedTemplate.Execute(w, data)
	}
}
