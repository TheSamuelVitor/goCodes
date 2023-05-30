package controllers

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func PaginaUsuarios(response http.ResponseWriter, request *http.Request) {

	templates = template.Must(template.ParseGlob("./pages/*.html"))

	templates.ExecuteTemplate(response, "index.html", nil)

}
