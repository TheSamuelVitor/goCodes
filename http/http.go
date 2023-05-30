package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {

	// rotas do servidor
	http.HandleFunc("/", home)
	http.HandleFunc("/usuarios", paginaUsuarios)

	// definicao do servidor
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func home(response http.ResponseWriter, request *http.Request) {

	// definicao de objeto
	type message struct {
		Message string `json:"mensagem"`
	}

	// criando um novo objeto
	ola := message{
		Message: "Pagina principal",
	}

	mensagem, error := json.Marshal(ola)
	if error != nil {
		log.Fatal(error)
	}

	response.Write(mensagem)
}

func paginaUsuarios(response http.ResponseWriter, request *http.Request) {

	templates = template.Must(template.ParseGlob("*.html"))

	templates.ExecuteTemplate(response, "index.html", nil)

}
