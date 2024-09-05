package main

import (
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		models.Personalidade{Nome: "Albert Einstein", Historia: "Físico alemão"},
		models.Personalidade{Nome: "Isaac Newton", Historia: "Físico e matemático inglês"},
		models.Personalidade{Nome: "Galileu Galilei", Historia: "Físico e astrônomo italiano"},
	}

	fmt.Println("Iniciando o servidor rest com go")
	routes.HandleRequests()
}
