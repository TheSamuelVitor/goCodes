package controllers

import (
	"http/models"
	"html/template"
	"net/http"
)

var templates *template.Template

func PaginaUsuarios(response http.ResponseWriter, request *http.Request) {

	usuario := model.Usuario{
		Name:  "Joao",
		Email: "joao@email.com",
	}

	templates = template.Must(template.ParseGlob("./view/*.html"))

	templates.ExecuteTemplate(response, "index.html", usuario)

}
