package main

import (
	"fmt"

	"github.com/kfsantos/livros/database"
	"github.com/kfsantos/livros/routes"
)

func main() {
	database.Conexao()
	fmt.Println("Iniciando o Servidor")
	routes.HandleRequest()
}
