package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Home(response http.ResponseWriter, request *http.Request) {

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
