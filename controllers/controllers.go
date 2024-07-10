package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WENDELLDELIMA/go-rest-api/database"
	"github.com/WENDELLDELIMA/go-rest-api/models"
	"github.com/gorilla/mux"
)


func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "home:")
}
func allPersonalidades() []models.Personalidade{
	var p []models.Personalidade
	database.DB.Find(&p)
	return p
}
func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	
	p := allPersonalidades()
	json.NewEncoder(w).Encode(p)
}
func RetornaUmapersonalidades(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.First(&p,id)
	if p.Id != 0 {
		json.NewEncoder(w).Encode(p)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Personalidade com id " + id + " não encontrada"})
	}
    
	
}

func CriaNovaPersonalidade(w http.ResponseWriter, r *http.Request){
	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)
	json.NewEncoder(w).Encode(personalidade)

}
func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    id := vars["id"]
    var p models.Personalidade
    database.DB.First(&p, id)
	if p.Id != 0 {
		database.DB.Delete(&p)
		json.NewEncoder(w).Encode(map[string]string{"message": "Personalidade de id " + id + " deletada"})
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Personalidade com id " + id + " não encontrada"})
	}
	
}
func EditaPersonalidade(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    id := vars["id"]
    var p models.Personalidade
	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}