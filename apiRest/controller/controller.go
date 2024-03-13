package controller

import (
	"apiRest/database"
	"apiRest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func PersonalidadePorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var pId []models.Personalidade

	database.DB.Find(&pId, id)

	json.NewEncoder(w).Encode(pId)

}

func AddPersonalidade(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adicionando:")
	var personalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&personalidade)

	fmt.Println(&personalidade)

	database.DB.Create(&personalidade)

	json.NewEncoder(w).Encode(personalidade)

}

func DeletePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}


func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.First(&personalidade,id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(&personalidade)

}

