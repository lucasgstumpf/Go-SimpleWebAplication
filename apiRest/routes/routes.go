package routes

import (
	"apiRest/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/personalidades", controller.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controller.PersonalidadePorId).Methods("Get")
	r.HandleFunc("/personalidades", controller.AddPersonalidade).Methods("Post")
	r.HandleFunc("/personalidades/{id}", controller.DeletePersonalidade).Methods("Delete")
	r.HandleFunc("/personalidades/{id}", controller.EditaPersonalidade).Methods("Put")
	http.ListenAndServe(":8000", r)

}
