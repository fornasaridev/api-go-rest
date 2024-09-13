package main

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Albert Einstein", Historia: "Físico alemão"},
		{Id: 2, Nome: "Isaac Newton", Historia: "Físico e matemático inglês"},
		{Id: 3, Nome: "Galileu Galilei", Historia: "Físico e astrônomo italiano"},
	}
	database.ConectaBancoDados()
	fmt.Println("Iniciando o servidor rest com go")
	routes.HandleResquest()
}
