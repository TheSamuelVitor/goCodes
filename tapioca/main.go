package main

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	
	var templates *template.Template

	templates = template.Must(template.ParseGlob("./pages/*.html"))

	server.GET("/", func(c *gin.Context) {

		c.JSON(200, templates.ExecuteTemplate(c.Writer, "index.html", nil))

	})

	fmt.Println("Servidor rodando na porta 3000")
	server.Run()
}
