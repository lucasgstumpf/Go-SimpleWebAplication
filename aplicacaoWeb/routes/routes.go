package routes

import (
	"aplicacaoWeb/controller"
	"net/http"
)

func BuscaRota() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)
	http.HandleFunc("/edit",controller.Edit)
	http.HandleFunc("/update", controller.Update)

}
