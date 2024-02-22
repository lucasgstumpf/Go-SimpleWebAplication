package main

import (
	"aplicacaoWeb/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.BuscaRota()
	http.ListenAndServe(":8000", nil)
}
