package routes

import (
	"apiRest/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/personalidades", controller.TodasPersonalidades)
	r.HandleFunc("/personalidades/{id}", controller.PersonalidadePorId)
	http.ListenAndServe(":8000", r)

}
