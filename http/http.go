package main

import (
	"log"
	"net/http"
	"http/controllers"
)

func main() {

	// rotas do servidor
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/usuarios", controllers.PaginaUsuarios)

	// definicao do servidor
	log.Fatal(http.ListenAndServe(":3000", nil))
}
