package main

import (
	"apiRest/models"
	"apiRest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 0, Nome: "nome 0", Descricao: "historia 0"},
		{Id: 1, Nome: "nome 1", Descricao: "historia 1"},
	}
	routes.HandleRequest()
}
