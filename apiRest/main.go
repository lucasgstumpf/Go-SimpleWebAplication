package main

import (
	"apiRest/database"
	"apiRest/routes"
)

func main() {

	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
