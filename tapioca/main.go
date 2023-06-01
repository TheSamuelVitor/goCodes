package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	fmt.Println("Servidor rodando na porta 3000")
	server.Run()
}
